package portfolio

import (
	"github.com/nekowawolf/airdropv2/utils"

	"github.com/google/uuid"
	"github.com/nekowawolf/airdropv2/config"

	"go.mongodb.org/mongo-driver/bson"
)

const portfolioCollection = "portfolio"

func GetPortfolio() (*Portfolio, error) {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	var result Portfolio
	err := config.Database.Collection(portfolioCollection).
		FindOne(ctx, bson.M{}).
		Decode(&result)
	return &result, err
}

func UpdatePortfolio(data Portfolio) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	_, err := config.Database.Collection(portfolioCollection).
		UpdateOne(ctx, bson.M{}, bson.M{"$set": data})
	return err
}

func UpdateHeroProfile(hero HeroProfile) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	_, err := config.Database.Collection(portfolioCollection).
		UpdateOne(
			ctx, 
			bson.M{}, 
			bson.M{"$set": bson.M{"hero": hero}},
		)
	return err
}

func addItem(field string, item interface{}) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()
	_, err := config.Database.Collection(portfolioCollection).
		UpdateOne(ctx, bson.M{}, bson.M{"$push": bson.M{field: item}})
	return err
}

func deleteItem(field, id string) error {
	ctx, cancel := utils.GetDBContext()
	defer cancel()

	_, err := config.Database.Collection(portfolioCollection).
		UpdateOne(ctx, bson.M{}, bson.M{"$pull": bson.M{field: bson.M{"id": id}}})
	return err
}

func AddExperience(e Experience) error {
	e.ID = uuid.NewString()
	return addItem("experience", e)
}

func AddEducation(e Education) error {
	e.ID = uuid.NewString()
	return addItem("education", e)
}

func AddTechSkill(s SkillItem) error {
	s.ID = uuid.NewString()
	return addItem("skills.tech", s)
}

func AddDesignSkill(s SkillItem) error {
	s.ID = uuid.NewString()
	return addItem("skills.design", s)
}

func DeleteExperience(id string) error {
	return deleteItem("experience", id)
}

func DeleteEducation(id string) error {
	return deleteItem("education", id)
}

func DeleteTechSkill(id string) error {
	return deleteItem("skills.tech", id)
}

func DeleteDesignSkill(id string) error {
	return deleteItem("skills.design", id)
}