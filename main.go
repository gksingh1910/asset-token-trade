package main

import "fmt"

type AssetToken struct {
	assetId        string
	assetName      string
	assetTokenCode string
}

type CreateAssetToken struct {
	Status string
	AssetToken
}

type AssetLifeCycle interface {
	createAsset(map[string]CreateAssetToken)
	launchAssetInOffering()
	changeAssetStatusInOffering()
	launchAssetInSm()
	changeAssetStatusInSm()
}

func (a *AssetToken) createAsset(asset map[string]CreateAssetToken) {
	asset[a.assetTokenCode] = CreateAssetToken{
		AssetToken: AssetToken{
			assetId:        a.assetId,
			assetName:      a.assetName,
			assetTokenCode: a.assetTokenCode,
		},
		Status: "WIP",
	}
	print("\n Asset token created")
	return
}

func (a *AssetToken) launchAssetInOffering() {

}
func (a *AssetToken) changeAssetStatusInOffering() {

}
func (a *AssetToken) launchAssetInSm() {

}
func (a *AssetToken) changeAssetStatusInSm() {

}

// func test(asCycle AssetLifeCycle) {
// 	b := make(map[string]CreateAssetToken)
// 	asCycle.createAsset(b)
// 	asCycle.launchAssetInOffering()
// 	asCycle.changeAssetStatusInSm()
// }

func main() {

	a := make(map[string]CreateAssetToken)
	AssetLifeCycle.createAsset(&AssetToken{
		assetId:        "1",
		assetName:      "Virat",
		assetTokenCode: "VKC",
	}, a)

	AssetLifeCycle.createAsset(&AssetToken{
		assetId:        "2",
		assetName:      "Rohit",
		assetTokenCode: "RHC",
	}, a)
	for key, value := range a {
		fmt.Printf("\n Map Key is  %v and Value is %v", key, value)
	}

}
