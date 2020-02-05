package types

// Validators moniker
var (
	Aconcagua     = "A71E5CD078B8C5C7B1AF88BCE84DD70B0557D93E"
	Ararat        = "B7707D9F593C62E85BB9E1A2366D12A97CD5DFF2"
	Carrauntoohil = "1175946A48EAA473868A0A6F52E6C66CCAF472EA"
	Everest       = "B0FBB52FF7EE93CC476DFE6B74FA1FC88584F30D"
	Elbrus        = "7235EF143D20FC0ABC427615D83014BB02D7C06C"
	Fuji          = "A9157B3FA6EB4C1E396B9B746E95327A07DC42E5"
	Gahinga       = "71F253E6FEA9EDD4B4753F5483549FE4F0F3A21C"
	Scafell       = "14CFCE69B645F3F88BAF08EA5B77FA521E4480F9"
	Seoraksan     = "17B42E8F284D3CA0E420262F89CD76C749BB12C9"
	Kita          = "414FB3BBA216AF84C47E07D6EBAA2DCFC3563A2F"
	Zugspitze     = "3CD4AABABDDEB7ABFEA9618732E331077A861D2B"
)

// GetValidatorMoniker returns validator's moniker (node name)
// after comparing their respective consensus hex address
func GetValidatorMoniker(address string) string {
	var moniker string

	switch {
	case address == Aconcagua:
		moniker = "Aconcagua"
	case address == Ararat:
		moniker = "Ararat"
	case address == Carrauntoohil:
		moniker = "Carrauntoohil"
	case address == Everest:
		moniker = "Everest"
	case address == Elbrus:
		moniker = "Elbrus"
	case address == Fuji:
		moniker = "Fuji"
	case address == Gahinga:
		moniker = "Gahinga"
	case address == Scafell:
		moniker = "Scafell"
	case address == Seoraksan:
		moniker = "Seoraksan"
	case address == Kita:
		moniker = "Kita"
	case address == Zugspitze:
		moniker = "Zugspitze"
	default:
		return ""
	}

	return moniker
}
