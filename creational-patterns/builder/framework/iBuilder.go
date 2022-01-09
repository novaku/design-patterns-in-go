package framework

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}

	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}