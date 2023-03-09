package elasticsearch

import(
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/valikhan03/auctions-service/models"
)

func InitElasticsearchConn() *elasticsearch.Client {

	esconfig := elasticsearch.Config{
		Addresses: models.ElasticConfigGlobal.Addrs,
		Username: models.ElasticConfigGlobal.Username,
	}
	
	client, err := elasticsearch.NewClient(esconfig)
	if err != nil{
		log.Fatalf("Elasticsearch connection error: %s", err.Error())
	}

	
	return client
}