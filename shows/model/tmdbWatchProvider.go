package model


type WatchProvider struct {
	ID      *int64   `json:"id,omitempty"`
	Results *Results `json:"results,omitempty"`
}

type Results struct {
	//Ar *Ar `json:"AR,omitempty"`
	//At *Ar `json:"AT,omitempty"`
	//Au *Ar `json:"AU,omitempty"`
	//Be *Ar `json:"BE,omitempty"`
	//Br *Ar `json:"BR,omitempty"`
	CA *Ar `json:"CA,omitempty"`
	//Ch *Ar `json:"CH,omitempty"`
	//Cl *Ar `json:"CL,omitempty"`
	//Co *Ar `json:"CO,omitempty"`
	//Cz *Ar `json:"CZ,omitempty"`
	//De *Ar `json:"DE,omitempty"`
	//Dk *Ar `json:"DK,omitempty"`
	//Ec *Ar `json:"EC,omitempty"`
	//Ee *Ar `json:"EE,omitempty"`
	//Es *Es `json:"ES,omitempty"`
	//Fi *Ar `json:"FI,omitempty"`
	//Fr *Ar `json:"FR,omitempty"`
	//GB *Ar `json:"GB,omitempty"`
	//Gr *Ar `json:"GR,omitempty"`
	//Hu *Ar `json:"HU,omitempty"`
	//ID *Ar `json:"ID,omitempty"`
	//Ie *Ar `json:"IE,omitempty"`
	//In *Ar `json:"IN,omitempty"`
	//It *Ar `json:"IT,omitempty"`
	//Jp *Ar `json:"JP,omitempty"`
	//Kr *Ar `json:"KR,omitempty"`
	//Lt *Ar `json:"LT,omitempty"`
	//LV *Ar `json:"LV,omitempty"`
	//MX *Ar `json:"MX,omitempty"`
	//My *Ar `json:"MY,omitempty"`
	//Nl *Ar `json:"NL,omitempty"`
	//No *Ar `json:"NO,omitempty"`
	//Nz *Ar `json:"NZ,omitempty"`
	//PE *Ar `json:"PE,omitempty"`
	//Ph *Ar `json:"PH,omitempty"`
	//Pl *Ar `json:"PL,omitempty"`
	//Pt *Ar `json:"PT,omitempty"`
	//Ro *Ro `json:"RO,omitempty"`
	//Ru *Ar `json:"RU,omitempty"`
	//SE *Ar `json:"SE,omitempty"`
	//Sg *Ar `json:"SG,omitempty"`
	//Th *Ar `json:"TH,omitempty"`
	//Tr *Ar `json:"TR,omitempty"`
	//Us *Ar `json:"US,omitempty"`
	//Ve *Ar `json:"VE,omitempty"`
	//Za *Es `json:"ZA,omitempty"`
}

type Ar struct {
	Link     *string `json:"link,omitempty"`
	Rent     []Ad    `json:"rent,omitempty"`
	Buy      []Ad    `json:"buy,omitempty"`
	Flatrate []Ad    `json:"flatrate,omitempty"`
	Ads      []Ad    `json:"ads,omitempty"`
}

type Ad struct {
	DisplayPriority *int64  `json:"display_priority,omitempty"`
	LogoPath        *string `json:"logo_path,omitempty"`
	ProviderID      *int64  `json:"provider_id,omitempty"`
	ProviderName    *string `json:"provider_name,omitempty"`
}

//type Es struct {
//	Link *string `json:"link,omitempty"`
//	Buy  []Ad    `json:"buy,omitempty"`
//}

//type Ro struct {
//	Link     *string `json:"link,omitempty"`
//	Flatrate []Ad    `json:"flatrate,omitempty"`
//}
