package main

// Collection struct
type Collection struct {
	RId          int       `json:"rid,omitempty"`
	Stac_version string    `json:"stac_version,omitempty"`
	Id           string    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	License      string    `json:"license,omitempty"`
	Keywords     []string  `json:"keywords,omitempty"`
	Providers    Providers `json:"providers,omitempty"`
	Extent       Extent    `json:"extent,omitempty"`
	Links        Links     `json:"links,omitempty"`
}

type Providers struct {
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Url   string   `json:"url,omitempty"`
}

type Extent struct {
	Spatial  []int    `json:"spatial,omitempty"`
	Temporal []string `json:"temporal,omitempty"`
}

// type Links struct {
// 	Href string `json:"href,omitempy"`
// 	Rel  string `json:"rel,omitempy"`
// }

type Collections []Collection

// {
// 	"stac_version": "0.6.1",
// 	"id": "COPERNICUS/S2",
// 	"title": "Sentinel-2 MSI: MultiSpectral Instrument, Level-1C",
// 	"description": "Sentinel-2 is a wide-swath, high-resolution, multi-spectral\nimaging mission supporting Copernicus Land Monitoring studies,\nincluding the monitoring of vegetation, soil and water cover,\nas well as observation of inland waterways and coastal areas.\n\nThe Sentinel-2 data contain 13 UINT16 spectral bands representing\nTOA reflectance scaled by 10000. See the [Sentinel-2 User Handbook](https://sentinel.esa.int/documents/247904/685211/Sentinel-2_User_Handbook)\nfor details. In addition, three QA bands are present where one\n(QA60) is a bitmask band with cloud mask information. For more\ndetails, [see the full explanation of how cloud masks are computed.](https://sentinel.esa.int/web/sentinel/technical-guides/sentinel-2-msi/level-1c/cloud-masks)\n\nEach Sentinel-2 product (zip archive) may contain multiple\ngranules. Each granule becomes a separate Earth Engine asset.\nEE asset ids for Sentinel-2 assets have the following format:\nCOPERNICUS/S2/20151128T002653_20151128T102149_T56MNN. Here the\nfirst numeric part represents the sensing date and time, the\nsecond numeric part represents the product generation date and\ntime, and the final 6-character string is a unique granule identifier\nindicating its UTM grid reference (see [MGRS](https://en.wikipedia.org/wiki/Military_Grid_Reference_System)).\n\nFor more details on Sentinel-2 radiometric resoltuon, [see this page](https://earth.esa.int/web/sentinel/user-guides/sentinel-2-msi/resolutions/radiometric).\n",
// 	"license": "proprietary",
// 	"keywords": [
// 	  "copernicus",
// 	  "esa",
// 	  "eu",
// 	  "msi",
// 	  "radiance",
// 	  "sentinel"
// 	],
// 	"providers": [
// 	  {
// 		"name": "European Union/ESA/Copernicus",
// 		"roles": [
// 		  "producer",
// 		  "licensor"
// 		],
// 		"url": "https://sentinel.esa.int/web/sentinel/user-guides/sentinel-2-msi"
// 	  }
// 	],
// 	"extent": {
// 	  "spatial": [
// 		-180.0,
// 		-56.0,
// 		180.0,
// 		83.0
// 	  ],
// 	  "temporal": [
// 		"2015-06-23T00:00:00Z",
// 		null
// 	  ]
// 	},
// 	"links": [
// 	  {
// 		"rel": "self",
// 		"href": "https://storage.cloud.google.com/earthengine-test/catalog/COPERNICUS_S2.json"
// 	  },
// 	  {
// 		"rel": "parent",
// 		"href": "https://storage.cloud.google.com/earthengine-test/catalog/catalog.json"
// 	  },
// 	  {
// 		"rel": "root",
// 		"href": "https://storage.cloud.google.com/earthengine-test/catalog/catalog.json"
// 	  },
// 	  {
// 		"rel": "license",
// 		"href": "https://scihub.copernicus.eu/twiki/pub/SciHubWebPortal/TermsConditions/Sentinel_Data_Terms_and_Conditions.pdf",
// 		"title": "Legal notice on the use of Copernicus Sentinel Data and Service Information"
// 	  }
// 	]
