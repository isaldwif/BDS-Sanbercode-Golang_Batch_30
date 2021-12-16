package main

func main() {
	var jsonString = `[{"vendor":"Acer","model":"beTouch E110"},{"vendor":"Acer","model":"beTouch E120"},{"vendor":"Acer","model":"beTouch E130"},{"vendor":"Acer","model":"beTouch E140"},{"vendor":"Acer","model":"beTouch E200"},{"vendor":"Acer","model":"beTouch E400"},{"vendor":"Acer","model":"Liquid"},{"vendor":"Acer","model":"Liquid E"},{"vendor":"Acer","model":"Liquid Metal"},{"vendor":"Acer","model":"neoTouch P300"},{"vendor":"Acer","model":"neoTouch P400"},{"vendor":"Acer","model":"Stream"},{"vendor":"Alcatel","model":"One Touch 701"},{"vendor":"Alcatel","model":"One Touch 715"},{"vendor":"Alcatel","model":"OT-980"},{"vendor":"Apple","model":"iPhone"},{"vendor":"Apple","model":"iPhone 3G"},{"vendor":"Apple","model":"iPhone 3GS"},{"vendor":"Apple","model":"iPhone 4"},{"vendor":"Apple","model":"iPhone 4S"},{"vendor":"Apple","model":"Iphone 5"},{"vendor":"BenQ-Siemens","model":"C81"},{"vendor":"BenQ-Siemens","model":"CF61"},{"vendor":"BenQ-Siemens","model":"CL71"},{"vendor":"BenQ-Siemens","model":"E61"},{"vendor":"BenQ-Siemens","model":"EL71"},{"vendor":"BenQ-Siemens","model":"P51"},{"vendor":"Dell","model":"Aero"},{"vendor":"Dell","model":"Streak"},{"vendor":"Dell","model":"Venue"},{"vendor":"Dell","model":"Venue Pro"},{"vendor":"Ericsson","model":"A3618s"},{"vendor":"Ericsson","model":"R380"},{"vendor":"Ericsson","model":"R520"},{"vendor":"Ericsson","model":"R600"},{"vendor":"Ericsson","model":"T29s"},{"vendor":"Ericsson","model":"T39"},{"vendor":"Ericsson","model":"T60d"},{"vendor":"Ericsson","model":"T65"},{"vendor":"Ericsson","model":"T66"},{"vendor":"Ericsson","model":"T68"},{"vendor":"Garmin-Asus","model":"nuvifone A50"},{"vendor":"Garmin-Asus","model":"nuvifone M10"},{"vendor":"HP","model":"Pre 3"},{"vendor":"HP","model":"Veer"},{"vendor":"HTC","model":"7 Mozart"},{"vendor":"HTC","model":"7 Pro"},{"vendor":"HTC","model":"7 Surround"},{"vendor":"HTC","model":"7 Trophy"},{"vendor":"HTC","model":"Advantage X7500"},{"vendor":"HTC","model":"Aria"},{"vendor":"HTC","model":"Arrive"},{"vendor":"HTC","model":"ChaCha"},{"vendor":"HTC","model":"Desire"},{"vendor":"HTC","model":"Desire C"},{"vendor":"HTC","model":"Desire C NFC"},{"vendor":"HTC","model":"Desire HD"},{"vendor":"HTC","model":"Desire S"},{"vendor":"HTC","model":"Desire VC"},{"vendor":"HTC","model":"Desire X"},{"vendor":"HTC","model":"Desire Z"},{"vendor":"HTC","model":"Droid Incredible"},{"vendor":"HTC","model":"EVO 3D"},{"vendor":"HTC","model":"EVO 4G"},{"vendor":"HTC","model":"EVO 4G LTE"},{"vendor":"HTC","model":"EVO Shift 4g"},{"vendor":"HTC","model":"Explorer"},{"vendor":"HTC","model":"Freestyle"},{"vendor":"HTC","model":"G1"},{"vendor":"HTC","model":"Google Nexus One"},{"vendor":"HTC","model":"Gratia"},{"vendor":"HTC","model":"HD Mini"},{"vendor":"HTC","model":"HD2"},{"vendor":"HTC","model":"HD7"},{"vendor":"HTC","model":"HD7S"},{"vendor":"HTC","model":"Hero"},{"vendor":"HTC","model":"Incredible S"},{"vendor":"HTC","model":"Inspire 4G"},{"vendor":"HTC","model":"J"},{"vendor":"HTC","model":"J Butterfly HTL21"},{"vendor":"HTC","model":"Legend"},{"vendor":"HTC","model":"Magic"},{"vendor":"HTC","model":"myTouch 3G Slide"},{"vendor":"HTC","model":"myTouch 4G"},{"vendor":"HTC","model":"One S"},{"vendor":"HTC","model":"One SV"},{"vendor":"HTC","model":"One V"},{"vendor":"HTC","model":"One X"},{"vendor":"HTC","model":"One X+"},{"vendor":"HTC","model":"Ozone"},{"vendor":"HTC","model":"P3300"},{"vendor":"HTC","model":"P3470"},{"vendor":"HTC","model":"P4350"},{"vendor":"HTC","model":"P6300"},{"vendor":"HTC","model":"P6500"},{"vendor":"HTC","model":"Radar"},{"vendor":"HTC","model":"Rhyme"},{"vendor":"HTC","model":"S310"},{"vendor":"HTC","model":"S620"},{"vendor":"HTC","model":"S710"},{"vendor":"HTC","model":"S730"},{"vendor":"HTC","model":"S740"},{"vendor":"HTC","model":"Sensation"},{"vendor":"HTC","model":"Sensation XE"},{"vendor":"HTC","model":"Sensation XL"},{"vendor":"HTC","model":"Shift"},{"vendor":"HTC","model":"Smart"},{"vendor":"HTC","model":"Snap"},{"vendor":"HTC","model":"Tattoo"},{"vendor":"HTC","model":"Thunderbolt"},{"vendor":"HTC","model":"Titan"},{"vendor":"HTC","model":"Titan II"},{"vendor":"HTC","model":"Touch"},{"vendor":"HTC","model":"Touch 3G"},{"vendor":"HTC","model":"Touch Cruise"},{"vendor":"HTC","model":"Touch Cruise 09"},{"vendor":"HTC","model":"Touch Diamond"},{"vendor":"HTC","model":"Touch Diamond 2"},{"vendor":"HTC","model":"Touch Dual"},{"vendor":"HTC","model":"Touch HD"},{"vendor":"HTC","model":"Touch Pro"},{"vendor":"HTC","model":"Touch Pro 2"},{"vendor":"HTC","model":"Touch Viva"},{"vendor":"HTC","model":"Touch2 T3333"},{"vendor":"HTC","model":"Trinity P3600"},{"vendor":"HTC","model":"TyTN II"},{"vendor":"HTC","model":"Wildfire"},{"vendor":"HTC","model":"Wildfire S"},{"vendor":"HTC","model":"Windows Phone 8S"},{"vendor":"HTC","model":"Windows Phone 8X"},{"vendor":"Huawei","model":"Ascend D quad"},{"vendor":"Huawei","model":"Ascend D1 Quad"},{"vendor":"Huawei","model":"Ascend G300 U8815"},{"vendor":"Huawei","model":"Ascend G600"},{"vendor":"Huawei","model":"Ascend P1"},{"vendor":"Kyocera","model":"Echo"},{"vendor":"LG","model":"7100"},{"vendor":"LG","model":"A120"},{"vendor":"LG","model":"A165"},{"vendor":"LG","model":"BL20 Chocolate"},{"vendor":"LG","model":"BL40 Chocolate"},{"vendor":"LG","model":"C105"},{"vendor":"LG","model":"C300"},{"vendor":"LG","model":"C310"},{"vendor":"LG","model":"Cookie Lite T300"},{"vendor":"LG","model":"Cookie Snap GM360i"},{"vendor":"LG","model":"Cosmos Touch VN270"},{"vendor":"LG","model":"enV Touch VX11000"},{"vendor":"LG","model":"eXpo GW820"},{"vendor":"LG","model":"F2300"},{"vendor":"LG","model":"G5300"},{"vendor":"LG","model":"G5400"},{"vendor":"LG","model":"G7000"},{"vendor":"LG","model":"GC900 Viewty Smart"},{"vendor":"LG","model":"GD510 Pop"},{"vendor":"LG","model":"GD900 Crystal"},{"vendor":"LG","model":"GS101"},{"vendor":"LG","model":"GS108"},{"vendor":"LG","model":"GT550 Encore"},{"vendor":"LG","model":"GW520"},{"vendor":"LG","model":"KC910 Renoir"},{"vendor":"LG","model":"KE970 Shine"},{"vendor":"LG","model":"KF510"},{"vendor":"LG","model":"KF600"},{"vendor":"LG","model":"KF700"},{"vendor":"LG","model":"KF750 Secret"},{"vendor":"LG","model":"KF900 Prada"},{"vendor":"LG","model":"KM900 Arena"},{"vendor":"LG","model":"KP500 Cookie"},{"vendor":"LG","model":"KS360"},{"vendor":"LG","model":"KU990 Viewty"},{"vendor":"LG","model":"M4410"},{"vendor":"LG","model":"Mini GD880"},{"vendor":"LG","model":"Nexus 4"},{"vendor":"LG","model":"Optimus 2X P990"},{"vendor":"LG","model":"Optimus 3D Max P720"},{"vendor":"LG","model":"Optimus 3D P920"},{"vendor":"LG","model":"Optimus 4X HD"},{"vendor":"LG","model":"Optimus 7 E900"},{"vendor":"LG","model":"Optimus 7Q - C900"},{"vendor":"LG","model":"Optimus Black P970"},{"vendor":"LG","model":"Optimus Chic E720"},{"vendor":"LG","model":"Optimus G"},{"vendor":"LG","model":"Optimus GT540"},{"vendor":"LG","model":"Optimus Hub E510"},{"vendor":"LG","model":"Optimus L3"},{"vendor":"LG","model":"Optimus L5"},{"vendor":"LG","model":"Optimus L5 Dual E612"},{"vendor":"LG","model":"Optimus L7 P700"},{"vendor":"LG","model":"Optimus L9"},{"vendor":"LG","model":"Optimus Me P350"},{"vendor":"LG","model":"Optimus Net P690"},{"vendor":"LG","model":"Optimus One P500"},{"vendor":"LG","model":"Optimus Pro C660"},{"vendor":"LG","model":"Optimus Sol E730"},{"vendor":"LG","model":"Optimus V"},{"vendor":"LG","model":"Optimus Vu"},{"vendor":"LG","model":"P520"},{"vendor":"LG","model":"P7200"},{"vendor":"LG","model":"Prada 3.0"},{"vendor":"LG","model":"Revolution"},{"vendor":"LG","model":"S310"},{"vendor":"LG","model":"Viewty GT400"},{"vendor":"LG","model":"Vu Plus"},{"vendor":"LG","model":"X330"},{"vendor":"LG","model":"Xenon GR500"},{"vendor":"Meizu","model":"Mx"},{"vendor":"Motorola","model":"A1000"},{"vendor":"Motorola","model":"A388c"},{"vendor":"Motorola","model":"A835"},{"vendor":"Motorola","model":"Admiral"},{"vendor":"Motorola","model":"Atrix 4G"},{"vendor":"Motorola","model":"Aura"},{"vendor":"Motorola","model":"Backflip"},{"vendor":"Motorola","model":"Barrage"},{"vendor":"Motorola","model":"Bravo"},{"vendor":"Motorola","model":"Brute i686"},{"vendor":"Motorola","model":"C350"},{"vendor":"Motorola","model":"Charm Motoblur"},{"vendor":"Motorola","model":"CLIQ 2"},{"vendor":"Motorola","model":"CLIQ MB200"},{"vendor":"Motorola","model":"Defy"},{"vendor":"Motorola","model":"Defy Mini"},{"vendor":"Motorola","model":"Defy Plus"},{"vendor":"Motorola","model":"Devour"},{"vendor":"Motorola","model":"DEXT MB200"},{"vendor":"Motorola","model":"Droid"},{"vendor":"Motorola","model":"Droid 2"},{"vendor":"Motorola","model":"Droid 2 Global"},{"vendor":"Motorola","model":"Droid 3"},{"vendor":"Motorola","model":"Droid 4"},{"vendor":"Motorola","model":"Droid Bionic"},{"vendor":"Motorola","model":"Droid Pro"},{"vendor":"Motorola","model":"Droid Razr HD"},{"vendor":"Motorola","model":"Droid Razr M"},{"vendor":"Motorola","model":"Droid Razr Maxx"},{"vendor":"Motorola","model":"Droid Razr Maxx HD"},{"vendor":"Motorola","model":"Droid Razr XT912"},{"vendor":"Motorola","model":"Droid X"},{"vendor":"Motorola","model":"Droid X2"},{"vendor":"Motorola","model":"E398"},{"vendor":"Motorola","model":"EX128"},{"vendor":"Motorola","model":"Flipout"},{"vendor":"Motorola","model":"Gleam Plus"},{"vendor":"Motorola","model":"Karma Qa1"},{"vendor":"Motorola","model":"Milestone"},{"vendor":"Motorola","model":"Milestone 2"},{"vendor":"Motorola","model":"Milestone XT720"},{"vendor":"Motorola","model":"MOTO RAZR V3"},{"vendor":"Motorola","model":"Motoluxe"},{"vendor":"Motorola","model":"Motoroi"},{"vendor":"Motorola","model":"Motosmart"},{"vendor":"Motorola","model":"MPx"},{"vendor":"Motorola","model":"MPx200"},{"vendor":"Motorola","model":"PEBL"},{"vendor":"Motorola","model":"Photon 4G"},{"vendor":"Motorola","model":"Pro+"},{"vendor":"Motorola","model":"Quench"},{"vendor":"Motorola","model":"Quench XT3 XT502"},{"vendor":"Motorola","model":"Quench XT5 XT502"},{"vendor":"Motorola","model":"RAZR"},
	{"vendor":"Motorola","model":"RAZR 2 V8"},{"vendor":"Motorola","model":"Razr i"},{"vendor":"Motorola","model":"RAZR MAXX"},{"vendor":"Motorola","model":"RAZR V3i"},{"vendor":"Motorola","model":"Rival A455"},{"vendor":"Motorola","model":"RIZR Z8"},{"vendor":"Motorola","model":"T720"},{"vendor":"Motorola","model":"Timeport 250"},{"vendor":"Motorola","model":"Timeport 260"},{"vendor":"Motorola","model":"V600"},{"vendor":"Motorola","model":"V70"},{"vendor":"Motorola","model":"V80"},{"vendor":"Motorola","model":"Wilder"},{"vendor":"Motorola","model":"XT531"},{"vendor":"Motorola","model":"ZN5 Motozine"},{"vendor":"Neonode","model":"N1"},{"vendor":"Nokia","model":"100"},{"vendor":"Nokia","model":"101"},{"vendor":"Nokia","model":"103"},{"vendor":"Nokia","model":"110"},{"vendor":"Nokia","model":"111"},{"vendor":"Nokia","model":"1110"},{"vendor":"Nokia","model":"1112"},{"vendor":"Nokia","model":"1200"},{"vendor":"Nokia","model":"1280"},{"vendor":"Nokia","model":"1616"},{"vendor":"Nokia","model":"1661"},{"vendor":"Nokia","model":"1662"},{"vendor":"Nokia","model":"1800"},{"vendor":"Nokia","model":"2100"},{"vendor":"Nokia","model":"2220 Slide"},{"vendor":"Nokia","model":"2310"},{"vendor":"Nokia","model":"2323 Classic"},{"vendor":"Nokia","model":"2330 Classic"},{"vendor":"Nokia","model":"2600 Classic"},{"vendor":"Nokia","model":"2610"},{"vendor":"Nokia","model":"2680 Slide"},{"vendor":"Nokia","model":"2690"},{"vendor":"Nokia","model":"2700 Classic"},{"vendor":"Nokia","model":"2710 Navigation Edition"},{"vendor":"Nokia","model":"2720 fold"},{"vendor":"Nokia","model":"2730 Classic"},{"vendor":"Nokia","model":"3100"},{"vendor":"Nokia","model":"3109 Classic"},{"vendor":"Nokia","model":"3110 Classic"},{"vendor":"Nokia","model":"3110 Evolve"},{"vendor":"Nokia","model":"3120 Classic"},{"vendor":"Nokia","model":"3200"},{"vendor":"Nokia","model":"3210"},{"vendor":"Nokia","model":"3220"},{"vendor":"Nokia","model":"3230"},{"vendor":"Nokia","model":"3250"},{"vendor":"Nokia","model":"3300"},{"vendor":"Nokia","model":"3300 Americas"},{"vendor":"Nokia","model":"3310"},{"vendor":"Nokia","model":"3330"},{"vendor":"Nokia","model":"3410"},{"vendor":"Nokia","model":"3500"},{"vendor":"Nokia","model":"3510"},{"vendor":"Nokia","model":"3510i"},{"vendor":"Nokia","model":"3530"},{"vendor":"Nokia","model":"3560"},{"vendor":"Nokia","model":"3585i"},{"vendor":"Nokia","model":"3600 Slide"},{"vendor":"Nokia","model":"3650"},{"vendor":"Nokia","model":"3660"},{"vendor":"Nokia","model":"3710 Fold"},{"vendor":"Nokia","model":"3720 Classic"},{"vendor":"Nokia","model":"500"},{"vendor":"Nokia","model":"5000"},{"vendor":"Nokia","model":"5030"},{"vendor":"Nokia","model":"5070"},{"vendor":"Nokia","model":"5100"},{"vendor":"Nokia","model":"5130 XpressMusic"},{"vendor":"Nokia","model":"5140"},{"vendor":"Nokia","model":"5200"},{"vendor":"Nokia","model":"5210"},{"vendor":"Nokia","model":"5220 XpressMusic"},{"vendor":"Nokia","model":"5228"},{"vendor":"Nokia","model":"5230"},{"vendor":"Nokia","model":"5235 Ovi Music Unlimited"},{"vendor":"Nokia","model":"5250"},{"vendor":"Nokia","model":"5300 XpressMusic"},{"vendor":"Nokia","model":"5310 XpressMusic"},{"vendor":"Nokia","model":"5330 Mobile TV Edition"},{"vendor":"Nokia","model":"5330 XpressMusic"},{"vendor":"Nokia","model":"5500"},{"vendor":"Nokia","model":"5510"},{"vendor":"Nokia","model":"5530 XpressMusic"},{"vendor":"Nokia","model":"5610 XpressMusic"},{"vendor":"Nokia","model":"5630 XpressMusic"},{"vendor":"Nokia","model":"5700 XpressMusic"},{"vendor":"Nokia","model":"5730 XpressMusic"},{"vendor":"Nokia","model":"5800 Navigation Edition"},{"vendor":"Nokia","model":"5800 XpressMusic"},{"vendor":"Nokia","model":"600"},{"vendor":"Nokia","model":"6021"},{"vendor":"Nokia","model":"603"},{"vendor":"Nokia","model":"6030"},{"vendor":"Nokia","model":"6060"},{"vendor":"Nokia","model":"6070"},{"vendor":"Nokia","model":"6100"},{"vendor":"Nokia","model":"6103"},{"vendor":"Nokia","model":"6110"},{"vendor":"Nokia","model":"6120 classic"},{"vendor":"Nokia","model":"6125"},{"vendor":"Nokia","model":"6131"},{"vendor":"Nokia","model":"6133"},{"vendor":"Nokia","model":"6136"},{"vendor":"Nokia","model":"6151"},{"vendor":"Nokia","model":"6200"},{"vendor":"Nokia","model":"6210"},{"vendor":"Nokia","model":"6210 Navigator"},{"vendor":"Nokia","model":"6212 Classic"},{"vendor":"Nokia","model":"6216 Classic"},{"vendor":"Nokia","model":"6220"},{"vendor":"Nokia","model":"6220 Classic"},{"vendor":"Nokia","model":"6230"},{"vendor":"Nokia","model":"6230i"},{"vendor":"Nokia","model":"6233"},{"vendor":"Nokia","model":"6250"},{"vendor":"Nokia","model":"6260"},{"vendor":"Nokia","model":"6267"},{"vendor":"Nokia","model":"6270"},{"vendor":"Nokia","model":"6280"},{"vendor":"Nokia","model":"6282"},{"vendor":"Nokia","model":"6288"},{"vendor":"Nokia","model":"6300"},{"vendor":"Nokia","model":"6301"},{"vendor":"Nokia","model":"6303 Classic"},{"vendor":"Nokia","model":"6303i Classic"},{"vendor":"Nokia","model":"6310"},{"vendor":"Nokia","model":"6500 Classic"},{"vendor":"Nokia","model":"6500 Slide"},{"vendor":"Nokia","model":"6510"},{"vendor":"Nokia","model":"6555"},{"vendor":"Nokia","model":"6600"},{"vendor":"Nokia","model":"6600 Fold"},{"vendor":"Nokia","model":"6600 Slide"},{"vendor":"Nokia","model":"6600i Slide"},{"vendor":"Nokia","model":"6610"},{"vendor":"Nokia","model":"6610i"},{"vendor":"Nokia","model":"6630"},{"vendor":"Nokia","model":"6650"},{"vendor":"Nokia","model":"6650 fold"},{"vendor":"Nokia","model":"6670"},{"vendor":"Nokia","model":"6680"},{"vendor":"Nokia","model":"6681"},{"vendor":"Nokia","model":"6700 Classic"},{"vendor":"Nokia","model":"6700 Slide"},{"vendor":"Nokia","model":"6710 Navigator"},{"vendor":"Nokia","model":"6720 Classic"},{"vendor":"Nokia","model":"6730 Classic"},{"vendor":"Nokia","model":"6760 Slide"},{"vendor":"Nokia","model":"6800"},{"vendor":"Nokia","model":"6820"},{"vendor":"Nokia","model":"700"},{"vendor":"Nokia","model":"701"},{"vendor":"Nokia","model":"7020"},{"vendor":"Nokia","model":"7070 Prism"},{"vendor":"Nokia","model":"7100 Supernova"},{"vendor":"Nokia","model":"7110"},{"vendor":"Nokia","model":"7200"},{"vendor":"Nokia","model":"7210"},{"vendor":"Nokia","model":"7210 Supernova"},{"vendor":"Nokia","model":"7230"},{"vendor":"Nokia","model":"7250"},{"vendor":"Nokia","model":"7250i"},{"vendor":"Nokia","model":"7260"},{"vendor":"Nokia","model":"7310 Supernova"},{"vendor":"Nokia","model":"7360"},{"vendor":"Nokia","model":"7370"},{"vendor":"Nokia","model":"7373"},{"vendor":"Nokia","model":"7380"},{"vendor":"Nokia","model":"7390"},{"vendor":"Nokia","model":"7500 Prism"},{"vendor":"Nokia","model":"7510 Supernova"},{"vendor":"Nokia","model":"7600"},{"vendor":"Nokia","model":"7610"},{"vendor":"Nokia","model":"7610 Supernova"},{"vendor":"Nokia","model":"7650"},{"vendor":"Nokia","model":"7700"},{"vendor":"Nokia","model":"7705 Twist"},{"vendor":"Nokia","model":"7710"},{"vendor":"Nokia","model":"7900 Prism"},{"vendor":"Nokia","model":"808 PureView"},{"vendor":"Nokia","model":"8210"},{"vendor":"Nokia","model":"8310"},{"vendor":"Nokia","model":"8600 Luna"},{"vendor":"Nokia","model":"8800"},{"vendor":"Nokia","model":"8800 Arte"},{"vendor":"Nokia","model":"8800 Sapphire Arte"},{"vendor":"Nokia","model":"8800 Sirocco"},{"vendor":"Nokia","model":"8850"},{"vendor":"Nokia","model":"8910"},{"vendor":"Nokia","model":"8910i"},{"vendor":"Nokia","model":"9210"},{"vendor":"Nokia","model":"9300"},{"vendor":"Nokia","model":"9300i"},{"vendor":"Nokia","model":"9500 Communicator"},{"vendor":"Nokia","model":"Asha 200"},{"vendor":"Nokia","model":"Asha 201"},{"vendor":"Nokia","model":"Asha 202"},{"vendor":"Nokia","model":"Asha 203"},{"vendor":"Nokia","model":"Asha 300"},{"vendor":"Nokia","model":"Asha 302"},{"vendor":"Nokia","model":"Asha 303"},{"vendor":"Nokia","model":"Asha 305"},{"vendor":"Nokia","model":"Asha 306"},{"vendor":"Nokia","model":"Asha 308"},{"vendor":"Nokia","model":"Asha 309"},{"vendor":"Nokia","model":"Asha 311"},{"vendor":"Nokia","model":"C1-00"},{"vendor":"Nokia","model":"C1-01"},{"vendor":"Nokia","model":"C1-02"},{"vendor":"Nokia","model":"C2-00"},{"vendor":"Nokia","model":"C2-01"},{"vendor":"Nokia","model":"C2-02"},{"vendor":"Nokia","model":"C2-03"},{"vendor":"Nokia","model":"C2-05"},{"vendor":"Nokia","model":"C2-06"},{"vendor":"Nokia","model":"C3"},{"vendor":"Nokia","model":"C3-01 Touch and Type"},{"vendor":"Nokia","model":"C5"},{"vendor":"Nokia","model":"C5-00 5MP"},{"vendor":"Nokia","model":"C5-03"},{"vendor":"Nokia","model":"C5-06"},{"vendor":"Nokia","model":"C6-00"},{"vendor":"Nokia","model":"C6-01"},{"vendor":"Nokia","model":"C7-00"},{"vendor":"Nokia","model":"E5"},{"vendor":"Nokia","model":"E50"},{"vendor":"Nokia","model":"E51"},{"vendor":"Nokia","model":"E52"},{"vendor":"Nokia","model":"E55"},{"vendor":"Nokia","model":"E6-00"},{"vendor":"Nokia","model":"E60"},{"vendor":"Nokia","model":"E61"},{"vendor":"Nokia","model":"E63"},{"vendor":"Nokia","model":"E65"},{"vendor":"Nokia","model":"E66"},{"vendor":"Nokia","model":"E7-00"},{"vendor":"Nokia","model":"E70"},{"vendor":"Nokia","model":"E71"},{"vendor":"Nokia","model":"E72"},{"vendor":"Nokia","model":"E75"},{"vendor":"Nokia","model":"E90"},{"vendor":"Nokia","model":"Lumia 510"},{"vendor":"Nokia","model":"Lumia 610"},{"vendor":"Nokia","model":"Lumia 610 NFC"},{"vendor":"Nokia","model":"Lumia 620"},{"vendor":"Nokia","model":"Lumia 710"},{"vendor":"Nokia","model":"Lumia 800"},{"vendor":"Nokia","model":"Lumia 820"},{"vendor":"Nokia","model":"Lumia 822"},{"vendor":"Nokia","model":"Lumia 900"},{"vendor":"Nokia","model":"Lumia 920"},{"vendor":"Nokia","model":"N-Gage"},{"vendor":"Nokia","model":"N70"},{"vendor":"Nokia","model":"N71"},{"vendor":"Nokia","model":"N72"},{"vendor":"Nokia","model":"N73"},{"vendor":"Nokia","model":"N75"},{"vendor":"Nokia","model":"N76"},{"vendor":"Nokia","model":"N77"},{"vendor":"Nokia","model":"N78"},{"vendor":"Nokia","model":"N79"},{"vendor":"Nokia","model":"N8"},{"vendor":"Nokia","model":"N80"},{"vendor":"Nokia","model":"N81"},
	{"vendor":"Nokia","model":"N81 8GB"},{"vendor":"Nokia","model":"N82"},{"vendor":"Nokia","model":"N85"},{"vendor":"Nokia","model":"N86 8MP"},{"vendor":"Nokia","model":"N9"},{"vendor":"Nokia","model":"N90"},{"vendor":"Nokia","model":"N900"},{"vendor":"Nokia","model":"N91"},{"vendor":"Nokia","model":"N92"},{"vendor":"Nokia","model":"N93"},{"vendor":"Nokia","model":"N95"},{"vendor":"Nokia","model":"N95 8GB"},{"vendor":"Nokia","model":"N96"},{"vendor":"Nokia","model":"N97"},{"vendor":"Nokia","model":"N97 Mini"},{"vendor":"Nokia","model":"Oro"},{"vendor":"Nokia","model":"Surge"},{"vendor":"Nokia","model":"X1-00"},{"vendor":"Nokia","model":"X1-01"},{"vendor":"Nokia","model":"X2"},{"vendor":"Nokia","model":"X2-01"},{"vendor":"Nokia","model":"X2-02"},{"vendor":"Nokia","model":"X2-05"},{"vendor":"Nokia","model":"X3"},{"vendor":"Nokia","model":"X3 Touch and Type"},{"vendor":"Nokia","model":"X5-01"},{"vendor":"Nokia","model":"X6 16GB"},{"vendor":"Nokia","model":"X6 32GB"},{"vendor":"Nokia","model":"X6 8GB"},{"vendor":"Nokia","model":"X7-00"},{"vendor":"Orange","model":"SPV"},{"vendor":"Palm","model":"Pixi"},{"vendor":"Palm","model":"Pixi Plus"},{"vendor":"Palm","model":"Pre"},{"vendor":"Palm","model":"Pre 2"},{"vendor":"Palm","model":"Pre Plus"},{"vendor":"Palm","model":"Treo 680"},{"vendor":"Palm","model":"Treo 750"},{"vendor":"Panasonic","model":"Eluga"},{"vendor":"Panasonic","model":"GD67"},{"vendor":"Panasonic","model":"GD87"},{"vendor":"Panasonic","model":"GD92"},{"vendor":"Panasonic","model":"X300"},{"vendor":"Panasonic","model":"X400"},{"vendor":"Panasonic","model":"X500"},{"vendor":"Philips","model":"Xenium 9@9"},{"vendor":"Philips","model":"Xenium X503"},{"vendor":"RIM","model":"Blackberry Bold 9000"},{"vendor":"RIM","model":"BlackBerry Bold 9700"},{"vendor":"RIM","model":"Blackberry Bold 9780"},{"vendor":"RIM","model":"Blackberry Bold 9790"},{"vendor":"RIM","model":"BlackBerry Bold 9900"},{"vendor":"RIM","model":"BlackBerry Bold 9930"},{"vendor":"RIM","model":"Blackberry Curve 3G 9300"},{"vendor":"RIM","model":"BlackBerry Curve 8300"},{"vendor":"RIM","model":"BlackBerry Curve 8310"},{"vendor":"RIM","model":"Blackberry Curve 8320"},{"vendor":"RIM","model":"Blackberry Curve 8520"},{"vendor":"RIM","model":"Blackberry Curve 8530"},{"vendor":"RIM","model":"Blackberry Curve 8900"},{"vendor":"RIM","model":"Blackberry Curve 9320"},{"vendor":"RIM","model":"BlackBerry Curve 9350"},{"vendor":"RIM","model":"BlackBerry Curve 9360"},{"vendor":"RIM","model":"BlackBerry Curve 9370"},{"vendor":"RIM","model":"Blackberry Curve 9380"},{"vendor":"RIM","model":"BlackBerry Pearl 3G 9100"},{"vendor":"RIM","model":"BlackBerry Pearl 3G 9105"},{"vendor":"RIM","model":"BlackBerry Porsche Design P9981"},{"vendor":"RIM","model":"BlackBerry Storm 9500"},{"vendor":"RIM","model":"BlackBerry Storm 9530"},{"vendor":"RIM","model":"BlackBerry Storm2 9520"},{"vendor":"RIM","model":"BlackBerry Storm2 9550"},{"vendor":"RIM","model":"Blackberry Style 9670"},{"vendor":"RIM","model":"Blackberry Torch 9800"},{"vendor":"RIM","model":"BlackBerry Torch 9810"},{"vendor":"RIM","model":"BlackBerry Tour 9630"},{"vendor":"Sagem","model":"myX-5"},{"vendor":"Samsung","model":"4G LTE"},{"vendor":"Samsung","model":"723 S7230"},{"vendor":"Samsung","model":"A257 Magnet"},{"vendor":"Samsung","model":"Ativ S"},{"vendor":"Samsung","model":"B2100"},{"vendor":"Samsung","model":"B2710"},{"vendor":"Samsung","model":"B3310"},{"vendor":"Samsung","model":"B7330"},{"vendor":"Samsung","model":"B7350 Omnia Pro 4"},{"vendor":"Samsung","model":"B7722"},{"vendor":"Samsung","model":"Beat DJ M7600"},{"vendor":"Samsung","model":"C3050"},{"vendor":"Samsung","model":"C3060"},{"vendor":"Samsung","model":"C3300 Champ"},{"vendor":"Samsung","model":"C3510 Genoa"},{"vendor":"Samsung","model":"C3530"},{"vendor":"Samsung","model":"C5010"},{"vendor":"Samsung","model":"C5130"},{"vendor":"Samsung","model":"C5212"},{"vendor":"Samsung","model":"C5510"},{"vendor":"Samsung","model":"C6625"},{"vendor":"Samsung","model":"Captivate i897"},{"vendor":"Samsung","model":"Ch@t 322 C3222"},{"vendor":"Samsung","model":"Ch@t 335 S3350"},{"vendor":"Samsung","model":"Champ Neo Duos GT-C3262"},{"vendor":"Samsung","model":"Corby Pro B5310"},{"vendor":"Samsung","model":"Corby S3650 Genio Touch"},{"vendor":"Samsung","model":"Diva S7070"},{"vendor":"Samsung","model":"E1080"},{"vendor":"Samsung","model":"E1120"},{"vendor":"Samsung","model":"E1360"},{"vendor":"Samsung","model":"E2120"},{"vendor":"Samsung","model":"E2210"},{"vendor":"Samsung","model":"E2370 xcover"},{"vendor":"Samsung","model":"E2550 Monte Slide"},{"vendor":"Samsung","model":"Epic 4G D700"},{"vendor":"Samsung","model":"Focus"},{"vendor":"Samsung","model":"Galaxy Ace 2"},{"vendor":"Samsung","model":"Galaxy Ace Plus"},{"vendor":"Samsung","model":"Galaxy Ace S5830"},{"vendor":"Samsung","model":"Galaxy Beam"},{"vendor":"Samsung","model":"Galaxy Fit S5670"},{"vendor":"Samsung","model":"Galaxy Gio S5660"},{"vendor":"Samsung","model":"Galaxy Grand GT-I9080"},{"vendor":"Samsung","model":"Galaxy Grand GT-I9082"},{"vendor":"Samsung","model":"Galaxy Mini 2"},{"vendor":"Samsung","model":"Galaxy Mini S5570"},{"vendor":"Samsung","model":"Galaxy Music"},{"vendor":"Samsung","model":"Galaxy Music Dual"},{"vendor":"Samsung","model":"Galaxy Nexus"},{"vendor":"Samsung","model":"Galaxy Note"},{"vendor":"Samsung","model":"Galaxy Note II"},{"vendor":"Samsung","model":"Galaxy Pocket"},{"vendor":"Samsung","model":"Galaxy Premier"},{"vendor":"Samsung","model":"Galaxy Pro"},{"vendor":"Samsung","model":"Galaxy R I9103"},{"vendor":"Samsung","model":"Galaxy S 4G"},{"vendor":"Samsung","model":"Galaxy S Advance"},{"vendor":"Samsung","model":"Galaxy S Duos S7562"},{"vendor":"Samsung","model":"Galaxy S i9000"},{"vendor":"Samsung","model":"Galaxy S II"},{"vendor":"Samsung","model":"Galaxy S II LTE"},{"vendor":"Samsung","model":"Galaxy S II WiMAX ISW11SC"},{"vendor":"Samsung","model":"Galaxy S III"},{"vendor":"Samsung","model":"Galaxy S III Mini"},{"vendor":"Samsung","model":"Galaxy SL I9003"},{"vendor":"Samsung","model":"Galaxy Spica I5700"},{"vendor":"Samsung","model":"Galaxy Tab P1000"},{"vendor":"Samsung","model":"Galaxy W"},{"vendor":"Samsung","model":"Galaxy Xcover GT-S5690"},{"vendor":"Samsung","model":"Galaxy Y Duos"},{"vendor":"Samsung","model":"Galaxy Y Pro"},{"vendor":"Samsung","model":"Galaxy Y Pro Duos"},{"vendor":"Samsung","model":"Gem I100"},{"vendor":"Samsung","model":"Google Nexus S I9020"},{"vendor":"Samsung","model":"I5500 Galaxy 5"},{"vendor":"Samsung","model":"I5800 Galaxy 3 Apollo"},{"vendor":"Samsung","model":"I7500 Galaxy"},{"vendor":"Samsung","model":"i8510 Innov8"},{"vendor":"Samsung","model":"Infuse 4G i997"},{"vendor":"Samsung","model":"Instinct HD"},{"vendor":"Samsung","model":"Intercept M910"},{"vendor":"Samsung","model":"Jet S8000"},{"vendor":"Samsung","model":"L700i"},{"vendor":"Samsung","model":"M2710"},{"vendor":"Samsung","model":"M5650 Lindy"},{"vendor":"Samsung","model":"Omnia 7"},{"vendor":"Samsung","model":"Omnia HD"},{"vendor":"Samsung","model":"Omnia i900"},{"vendor":"Samsung","model":"Omnia II I8000"},{"vendor":"Samsung","model":"Omnia Pro B7610"},{"vendor":"Samsung","model":"Omnia W"},{"vendor":"Samsung","model":"Pixon12 M8910"},{"vendor":"Samsung","model":"Rugby II A847"},{"vendor":"Samsung","model":"S5600"},{"vendor":"Samsung","model":"S5620 Monte"},{"vendor":"Samsung","model":"S7550 Blue Earth"},{"vendor":"Samsung","model":"SGH-A800"},{"vendor":"Samsung","model":"SGH-C100"},{"vendor":"Samsung","model":"SGH-D500"},{"vendor":"Samsung","model":"SGH-D600"},{"vendor":"Samsung","model":"SGH-D600E"},{"vendor":"Samsung","model":"SGH-D840"},{"vendor":"Samsung","model":"SGH-D900"},{"vendor":"Samsung","model":"SGH-E250"},{"vendor":"Samsung","model":"SGH-E700"},{"vendor":"Samsung","model":"SGH-E730"},{"vendor":"Samsung","model":"SGH-E800"},{"vendor":"Samsung","model":"SGH-F300"},{"vendor":"Samsung","model":"SGH-F330"},{"vendor":"Samsung","model":"SGH-G600"},{"vendor":"Samsung","model":"SGH-G800"},{"vendor":"Samsung","model":"SGH-i320"},{"vendor":"Samsung","model":"SGH-i600"},{"vendor":"Samsung","model":"SGH-M300"},{"vendor":"Samsung","model":"SGH-P730"},{"vendor":"Samsung","model":"SGH-S100"},{"vendor":"Samsung","model":"SGH-S300"},{"vendor":"Samsung","model":"SGH-T100"},{"vendor":"Samsung","model":"SGH-U600"},{"vendor":"Samsung","model":"SGH-U700"},{"vendor":"Samsung","model":"SGH-X100"},{"vendor":"Samsung","model":"SGH-X450"},{"vendor":"Samsung","model":"SGH-X640"},{"vendor":"Samsung","model":"SGH-Z500"},{"vendor":"Samsung","model":"Shark S5350"},{"vendor":"Samsung","model":"Showcase i500"},{"vendor":"Samsung","model":"Star 3"},{"vendor":"Samsung","model":"Star 3 DUOS"},{"vendor":"Samsung","model":"Star II S5260"},{"vendor":"Samsung","model":"Strive A687"},{"vendor":"Samsung","model":"Sunburst A697"},{"vendor":"Samsung","model":"T-mobile Sidekick 4G"},{"vendor":"Samsung","model":"Tocco Lite S5230"},{"vendor":"Samsung","model":"Transform M920"},{"vendor":"Samsung","model":"U900 Soul"},{"vendor":"Samsung","model":"Ultra B S7220"},{"vendor":"Samsung","model":"Ultra S S7350"},{"vendor":"Samsung","model":"Ultra Touch S8300"},{"vendor":"Samsung","model":"Wave 3"},{"vendor":"Samsung","model":"Wave 575 S5750"},{"vendor":"Samsung","model":"Wave II S8530"},{"vendor":"Samsung","model":"Wave M"},{"vendor":"Samsung","model":"Wave S8500"},{"vendor":"Samsung","model":"Wave Y S5380"},{"vendor":"Sendo","model":"X"},{"vendor":"Sharp","model":"Aquos 007SH"},{"vendor":"Sharp","model":"GX10"},{"vendor":"Siemens","model":"A55"},{"vendor":"Siemens","model":"C60"},{"vendor":"Siemens","model":"C62"},{"vendor":"Siemens","model":"C65"},{"vendor":"Siemens","model":"C75"},{"vendor":"Siemens","model":"CX65"},{"vendor":"Siemens","model":"M50"},{"vendor":"Siemens","model":"M65"},{"vendor":"Siemens","model":"S40"},{"vendor":"Siemens","model":"S55"},{"vendor":"Siemens","model":"S65"},{"vendor":"Siemens","model":"SF65"},{"vendor":"Siemens","model":"SK65"},
	{"vendor":"Siemens","model":"SL42"},{"vendor":"Siemens","model":"SL45i"},{"vendor":"Siemens","model":"SL55"},{"vendor":"Siemens","model":"SX1"},{"vendor":"Sony","model":"CMD-Z5"},{"vendor":"Sony","model":"Xperia Acro S"},{"vendor":"Sony","model":"Xperia Advance"},{"vendor":"Sony","model":"Xperia E"},{"vendor":"Sony","model":"Xperia E Dual"},{"vendor":"Sony","model":"Xperia Go"},{"vendor":"Sony","model":"Xperia Ion LT28at"},{"vendor":"Sony","model":"Xperia J"},{"vendor":"Sony","model":"Xperia Miro"},{"vendor":"Sony","model":"Xperia Neo L"},{"vendor":"Sony","model":"Xperia P"},{"vendor":"Sony","model":"Xperia S"},{"vendor":"Sony","model":"Xperia SL"},{"vendor":"Sony","model":"Xperia Sola"},{"vendor":"Sony","model":"Xperia SX"},{"vendor":"Sony","model":"Xperia T"},{"vendor":"Sony","model":"Xperia Tipo"},{"vendor":"Sony","model":"Xperia Tipo Dual"},{"vendor":"Sony","model":"Xperia TL"},{"vendor":"Sony","model":"Xperia TX"},{"vendor":"Sony","model":"Xperia U"},{"vendor":"Sony","model":"Xperia V"},{"vendor":"Sony","model":"Xperia VL"},{"vendor":"Sony","model":"Ericsson Aino"},{"vendor":"Sony","model":"Ericsson Aspen"},{"vendor":"Sony","model":"Ericsson C510"},{"vendor":"Sony","model":"Ericsson C702"},{"vendor":"Sony","model":"Ericsson C901"},{"vendor":"Sony","model":"Ericsson C902"},{"vendor":"Sony","model":"Ericsson C903"},{"vendor":"Sony","model":"Ericsson C905"},{"vendor":"Sony","model":"Ericsson Cedar J108i"},{"vendor":"Sony","model":"Ericsson D750i"},{"vendor":"Sony","model":"Ericsson Elm J10"},{"vendor":"Sony","model":"Ericsson Elm J10i2"},{"vendor":"Sony","model":"Ericsson Equinox T717"},{"vendor":"Sony","model":"Ericsson F305"},{"vendor":"Sony","model":"Ericsson F500i"},{"vendor":"Sony","model":"Ericsson G502"},{"vendor":"Sony","model":"Ericsson G700"},{"vendor":"Sony","model":"Ericsson G705"},{"vendor":"Sony","model":"Ericsson G900"},{"vendor":"Sony","model":"Ericsson Hazel J20"},{"vendor":"Sony","model":"Ericsson Hazel J20i"},{"vendor":"Sony","model":"Ericsson J100"},{"vendor":"Sony","model":"Ericsson J110"},{"vendor":"Sony","model":"Ericsson J120"},{"vendor":"Sony","model":"Ericsson J132"},{"vendor":"Sony","model":"Ericsson J200i"},{"vendor":"Sony","model":"Ericsson J210"},{"vendor":"Sony","model":"Ericsson J220"},{"vendor":"Sony","model":"Ericsson J230"},{"vendor":"Sony","model":"Ericsson J300i"},{"vendor":"Sony","model":"Ericsson Jalou"},{"vendor":"Sony","model":"Ericsson Jalou Dolce Gabbana"},{"vendor":"Sony","model":"Ericsson K200"},{"vendor":"Sony","model":"Ericsson K220"},{"vendor":"Sony","model":"Ericsson K300"},{"vendor":"Sony","model":"Ericsson K310"},{"vendor":"Sony","model":"Ericsson K320"},{"vendor":"Sony","model":"Ericsson K330"},{"vendor":"Sony","model":"Ericsson K500"},{"vendor":"Sony","model":"Ericsson K510"},{"vendor":"Sony","model":"Ericsson K530"},{"vendor":"Sony","model":"Ericsson K550"},{"vendor":"Sony","model":"Ericsson K600"},{"vendor":"Sony","model":"Ericsson K608"},{"vendor":"Sony","model":"Ericsson K610"},{"vendor":"Sony","model":"Ericsson K618"},{"vendor":"Sony","model":"Ericsson K630"},{"vendor":"Sony","model":"Ericsson K660"},{"vendor":"Sony","model":"Ericsson K700"},{"vendor":"Sony","model":"Ericsson K750"},{"vendor":"Sony","model":"Ericsson K770"},{"vendor":"Sony","model":"Ericsson K790"},{"vendor":"Sony","model":"Ericsson K800"},{"vendor":"Sony","model":"Ericsson K810"},{"vendor":"Sony","model":"Ericsson K850"},{"vendor":"Sony","model":"Ericsson Kita"},{"vendor":"Sony","model":"Ericsson Live Walkman"},{"vendor":"Sony","model":"Ericsson M600"},{"vendor":"Sony","model":"Ericsson Mix Walkman"},{"vendor":"Sony","model":"Ericsson Naite"},{"vendor":"Sony","model":"Ericsson P1"},{"vendor":"Sony","model":"Ericsson P800"},{"vendor":"Sony","model":"Ericsson P900"},{"vendor":"Sony","model":"Ericsson P910"},{"vendor":"Sony","model":"Ericsson P910a"},{"vendor":"Sony","model":"Ericsson P990"},{"vendor":"Sony","model":"Ericsson R300"},{"vendor":"Sony","model":"Ericsson R306"},{"vendor":"Sony","model":"Ericsson R306a"},{"vendor":"Sony","model":"Ericsson S302"},{"vendor":"Sony","model":"Ericsson S312"},{"vendor":"Sony","model":"Ericsson S500"},{"vendor":"Sony","model":"Ericsson S600"},{"vendor":"Sony","model":"Ericsson S700i"},{"vendor":"Sony","model":"Ericsson S710a"},{"vendor":"Sony","model":"Ericsson Satio"},{"vendor":"Sony","model":"Ericsson Spiro"},{"vendor":"Sony","model":"Ericsson T100"},{"vendor":"Sony","model":"Ericsson T200"},{"vendor":"Sony","model":"Ericsson T226"},{"vendor":"Sony","model":"Ericsson T230"},{"vendor":"Sony","model":"Ericsson T250"},{"vendor":"Sony","model":"Ericsson T270"},{"vendor":"Sony","model":"Ericsson T280"},{"vendor":"Sony","model":"Ericsson T280a"},{"vendor":"Sony","model":"Ericsson T290"},{"vendor":"Sony","model":"Ericsson T300"},{"vendor":"Sony","model":"Ericsson T303"},{"vendor":"Sony","model":"Ericsson T306"},{"vendor":"Sony","model":"Ericsson T310"},{"vendor":"Sony","model":"Ericsson T316"},{"vendor":"Sony","model":"Ericsson T600"},{"vendor":"Sony","model":"Ericsson T606"},{"vendor":"Sony","model":"Ericsson T608"},{"vendor":"Sony","model":"Ericsson T610"},{"vendor":"Sony","model":"Ericsson T616"},{"vendor":"Sony","model":"Ericsson T618"},{"vendor":"Sony","model":"Ericsson T630"},{"vendor":"Sony","model":"Ericsson T650"},{"vendor":"Sony","model":"Ericsson T658"},{"vendor":"Sony","model":"Ericsson T68i"},{"vendor":"Sony","model":"Ericsson T700"},{"vendor":"Sony","model":"Ericsson T707"},{"vendor":"Sony","model":"Ericsson T715"},{"vendor":"Sony","model":"Ericsson Txt"},{"vendor":"Sony","model":"Ericsson Txt Pro"},{"vendor":"Sony","model":"Ericsson V600"},{"vendor":"Sony","model":"Ericsson V800"},{"vendor":"Sony","model":"Ericsson Vivaz"},{"vendor":"Sony","model":"Ericsson Vivaz Pro"},{"vendor":"Sony","model":"Ericsson W200"},{"vendor":"Sony","model":"Ericsson W205"},{"vendor":"Sony","model":"Ericsson W205a"},{"vendor":"Sony","model":"Ericsson W300"},{"vendor":"Sony","model":"Ericsson W302"},{"vendor":"Sony","model":"Ericsson W350"},{"vendor":"Sony","model":"Ericsson W350a"},{"vendor":"Sony","model":"Ericsson W380"},{"vendor":"Sony","model":"Ericsson W395"},{"vendor":"Sony","model":"Ericsson W508"},{"vendor":"Sony","model":"Ericsson W518a"},{"vendor":"Sony","model":"Ericsson W550"},{"vendor":"Sony","model":"Ericsson W580"},{"vendor":"Sony","model":"Ericsson W595"},{"vendor":"Sony","model":"Ericsson W600"},{"vendor":"Sony","model":"Ericsson W610"},{"vendor":"Sony","model":"Ericsson W660"},{"vendor":"Sony","model":"Ericsson W700"},{"vendor":"Sony","model":"Ericsson W705"},{"vendor":"Sony","model":"Ericsson W710"},{"vendor":"Sony","model":"Ericsson W715"},{"vendor":"Sony","model":"Ericsson W760"},{"vendor":"Sony","model":"Ericsson W8 Walkman"},{"vendor":"Sony","model":"Ericsson W800"},{"vendor":"Sony","model":"Ericsson W810"},{"vendor":"Sony","model":"Ericsson W830"},{"vendor":"Sony","model":"Ericsson W850"},{"vendor":"Sony","model":"Ericsson W880"},{"vendor":"Sony","model":"Ericsson W890"},{"vendor":"Sony","model":"Ericsson W900"},{"vendor":"Sony","model":"Ericsson W902"},{"vendor":"Sony","model":"Ericsson W910"},{"vendor":"Sony","model":"Ericsson W950"},{"vendor":"Sony","model":"Ericsson W960"},{"vendor":"Sony","model":"Ericsson W980"},{"vendor":"Sony","model":"Ericsson W995"},{"vendor":"Sony","model":"Ericsson Xperia Active"},{"vendor":"Sony","model":"Ericsson Xperia Arc"},{"vendor":"Sony","model":"Ericsson Xperia Arc S"},{"vendor":"Sony","model":"Ericsson Xperia Mini"},{"vendor":"Sony","model":"Ericsson Xperia Mini Pro"},{"vendor":"Sony","model":"Ericsson Xperia Neo"},{"vendor":"Sony","model":"Ericsson Xperia Neo V"},{"vendor":"Sony","model":"Ericsson Xperia Play"},{"vendor":"Sony","model":"Ericsson Xperia Play 4G"},{"vendor":"Sony","model":"Ericsson Xperia Pro"},{"vendor":"Sony","model":"Ericsson Xperia Pureness"},{"vendor":"Sony","model":"Ericsson Xperia Ray"},{"vendor":"Sony","model":"Ericsson Xperia X1"},{"vendor":"Sony","model":"Ericsson Xperia X10"},{"vendor":"Sony","model":"Ericsson Xperia X10 Mini"},{"vendor":"Sony","model":"Ericsson Xperia X10 Mini Pro"},{"vendor":"Sony","model":"Ericsson Xperia X2"},{"vendor":"Sony","model":"Ericsson Xperia X8"},{"vendor":"Sony","model":"Ericsson Yari"},{"vendor":"Sony","model":"Ericsson Yendo"},{"vendor":"Sony","model":"Ericsson Z1010"},{"vendor":"Sony","model":"Ericsson Z200"},{"vendor":"Sony","model":"Ericsson Z250"},{"vendor":"Sony","model":"Ericsson Z300"},{"vendor":"Sony","model":"Ericsson Z310"},{"vendor":"Sony","model":"Ericsson Z320"},{"vendor":"Sony","model":"Ericsson Z500"},{"vendor":"Sony","model":"Ericsson Z520"},{"vendor":"Sony","model":"Ericsson Z525"},{"vendor":"Sony","model":"Ericsson Z530"},{"vendor":"Sony","model":"Ericsson Z550"},{"vendor":"Sony","model":"Ericsson Z555"},{"vendor":"Sony","model":"Ericsson Z558"},{"vendor":"Sony","model":"Ericsson Z600"},{"vendor":"Sony","model":"Ericsson Z610"},{"vendor":"Sony","model":"Ericsson Z710"},{"vendor":"Sony","model":"Ericsson Z750"},{"vendor":"Sony","model":"Ericsson Z770"},{"vendor":"Sony","model":"Ericsson Z780"},{"vendor":"Sony","model":"Ericsson Z800i"},{"vendor":"Sony","model":"Ericsson Zylo"},{"vendor":"Tel.Me.","model":"T919"},{"vendor":"Toshiba","model":"G450"},{"vendor":"Toshiba","model":"Portege G500"},{"vendor":"Toshiba","model":"Portege G710"},{"vendor":"Toshiba","model":"Portege G810"},{"vendor":"Toshiba","model":"Portege G910"},{"vendor":"Toshiba","model":"TG01"},{"vendor":"Toshiba","model":"TS21i"},{"vendor":"ViewSonic","model":"ViewPad 4"},{"vendor":"ViewSonic","model":"ViewPhone 3"}]`
}
