package constants

const (
	LotDataOverTimeFull            = LotDataOverTimeNoID + "/{lotID}"
	LotDataOverTimeNoID            = "/v1/lotdataovertime"
	LotByID                        = Lots + "/{lotID}"
	Lots                           = "/v1/lot"
	UntrackedLots                  = "/v1/untrackedlot"
	UntrackedLotsByID              = UntrackedLots + "/{lotID}"
	Permits                  = "/v1/permit"
	PermitsByID              = Permits + "/{permitID}"
	LotAverageFreespaceByDayNoDate = "/v1/avgfree"
	LotAverageFreespaceByDay       = LotAverageFreespaceByDayNoDate // + "/{day}/{time}"

	RegisterAppInstance = "/v1/register"
)
