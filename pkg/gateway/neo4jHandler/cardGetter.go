package neo4jHandler

import (
	"fmt"

	"github.com/SeoSoojin/TKG_Kards/pkg/domain/models"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jGetter struct {
	Driver  *neo4j.Driver
	Session *neo4j.Session
}

func NewNeo4jGetter(driver *neo4j.Driver) *Neo4jGetter {

	if driver == nil {

		panic("driver is nil")

	}
	driverAux := *driver
	session := driverAux.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	return &Neo4jGetter{
		Driver:  driver,
		Session: &session,
	}

}

func (getter *Neo4jGetter) GetCardsByUserId(userId int) ([]models.Card, error) {

	query := `MATCH (u:User)-[h:HAS]->(c:Card)-[:PART_OF]->(col:Collection) 
			WHERE u.discordId = $discordId
			RETURN c, h, col`

	userIdString := fmt.Sprintf("%d", userId)
	queryArgs := models.QueryArgs{
		Key:   "discordId",
		Value: userIdString,
	}
	session := *(getter.Session)
	cards, err := session.ReadTransaction(fetchCardsInv(query, queryArgs))

	if err != nil {

		return nil, err

	}

	return cards.([]models.Card), nil

}

func (getter *Neo4jGetter) GetCardsByCollectionId(collectionId int) ([]models.Card, error) {

	query := `MATCH (c:Card)-[:PART_OF]->(col:Collection) 
			WHERE id(col) = $collectionId
			RETURN c, col`
	queryArgs := models.QueryArgs{
		Key:   "collectionId",
		Value: collectionId,
	}
	session := *(getter.Session)
	cards, err := session.ReadTransaction(fetchCards(query, queryArgs))
	if err != nil {

		return nil, err

	}
	return cards.([]models.Card), nil

}

func (getter *Neo4jGetter) GetCardsByIdolId(idolId int) ([]models.Card, error) {

	session := *(getter.Session)
	query := `MATCH (i:Idol)-[:IN]->(c:Card)-[:PART_OF]->(col:Collection)
			WHERE id(i) = $idolId
			RETURN c, col`
	queryArgs := models.QueryArgs{
		Key:   "idolId",
		Value: idolId,
	}
	cards, err := session.ReadTransaction(fetchCards(query, queryArgs))
	if err != nil {

		return nil, err

	}
	return cards.([]models.Card), nil

}

func (getter *Neo4jGetter) GetCardsByGroupId(groupId int) ([]models.Card, error) {

	session := *(getter.Session)
	query := `MATCH (g:Group)-[:RELATED_TO]->(col:Collection)<-[:PART_OF]->(c:Card)
			WHERE id(g) = $groupId
			RETURN c, col`
	queryArgs := models.QueryArgs{
		Key:   "groupId",
		Value: groupId,
	}
	cards, err := session.ReadTransaction(fetchCards(query, queryArgs))
	if err != nil {

		return nil, err

	}
	return cards.([]models.Card), nil

}

func (getter *Neo4jGetter) GetCardsByAlbumId(albumId int) ([]models.Card, error) {

	session := *(getter.Session)
	query := `MATCH (a:Album)-[:RELATED_TO]->(col:Collection)<-[:PART_OF]->(c:Card)
				WHERE id(a) = $albumId
				RETURN c, col`
	queryArgs := models.QueryArgs{
		Key:   "albumId",
		Value: albumId,
	}
	cards, err := session.ReadTransaction(fetchCards(query, queryArgs))
	if err != nil {

		return nil, err

	}
	return cards.([]models.Card), nil

}

func (getter *Neo4jGetter) GetCards() ([]models.Card, error) {

	session := *(getter.Session)
	query := `MATCH (c:Card)-[:PART_OF]->(col:Collection) 
			RETURN c, col`
	cards, err := session.ReadTransaction(fetchCards(query))
	if err != nil {

		return nil, err

	}
	return cards.([]models.Card), nil

}

func fetchCards(query string, queryArgs ...models.QueryArgs) neo4j.TransactionWork {

	mapped := make(map[string]interface{})

	for _, arg := range queryArgs {

		mapped[arg.Key] = arg.Value

	}
	return func(tx neo4j.Transaction) (interface{}, error) {

		records, err := tx.Run(query, mapped)
		if err != nil {

			fmt.Println(err)
			return nil, err

		}
		var cards []models.Card
		for records.Next() {
			record := records.Record()
			cardNode := record.Values[0].(neo4j.Node)
			colNode := record.Values[1].(neo4j.Node)
			card := models.Card{
				Id:          int(cardNode.Id),
				Name:        cardNode.Props["name"].(string),
				Rarity:      int(cardNode.Props["rarity"].(int64)),
				Type:        cardNode.Props["type"].(string),
				ImgRef:      cardNode.Props["imgRef"].(string),
				ReleaseDate: cardNode.Props["releaseDate"].(string),
				Collection:  colNode.Props["name"].(string),
			}
			cards = append(cards, card)
		}
		return cards, nil
	}
}

func fetchCardsInv(query string, queryArgs ...models.QueryArgs) neo4j.TransactionWork {

	mapped := make(map[string]interface{})

	for _, arg := range queryArgs {

		mapped[arg.Key] = arg.Value

	}
	return func(tx neo4j.Transaction) (interface{}, error) {

		records, err := tx.Run(query, mapped)
		if err != nil {

			fmt.Println(err)
			return nil, err

		}
		var cards []models.Card
		for records.Next() {
			record := records.Record()
			cardNode := record.Values[0].(neo4j.Node)
			hasNode := record.Values[1].(neo4j.Relationship)
			colNode := record.Values[2].(neo4j.Node)
			card := models.Card{
				Id:          int(cardNode.Id),
				Name:        cardNode.Props["name"].(string),
				Rarity:      int(cardNode.Props["rarity"].(int64)),
				Type:        cardNode.Props["type"].(string),
				ImgRef:      cardNode.Props["imgRef"].(string),
				ReleaseDate: cardNode.Props["releaseDate"].(string),
				Collection:  colNode.Props["name"].(string),
				Qnt:         int(hasNode.Props["qnt"].(int64)),
				Favorite:    hasNode.Props["fav"].(bool),
			}
			cards = append(cards, card)
		}
		return cards, nil
	}
}
