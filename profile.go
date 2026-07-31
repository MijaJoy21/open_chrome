package main

type ChromeProfile struct {
	Name string
	URLs []string
}

// 	"Profile 2", // profile xs kins  //
// 	//"Profile 4",  // profile kins ungu //
// 	//"Profile 5",  // profile monkey tech //
// 	//"Profile 19", // profile smart cuy //
// 	//"Profile 10", // profile mija //
// 	//"Profile 21", // profile shopee ajim //
// 	//"Profile 7", // profile unclemi //
// 	//"Profile 8", // profile xmilabs //
// 	//"Profile 9", // profile 360 tech //
// 	//"Profile 12", // profile adm gaje //
// 	//"Profile 13", // profile akun univ ajim //
// 	//"Profile 15", // profile akun ajim gaje //
// 	//"Profile 17", // profile adm sl kins //

var profiles = []ChromeProfile{
	{
		Name: "Profile 12", // profile xs kins  //
		URLs: []string{
			"https://account.accurate.id/manage",
		},
	},
	{
		Name: "Profile 8", // profile kins ungu //
		URLs: []string{
			" https://subaccount.shopee.com/chat-distribution/shops",
		},
	},
	{
		Name: "Profile 17", // profile monkey tech //
		URLs: []string{
			"https://seller-id.tokopedia.com/order?order_status[]=1&selected_sort=11&tab=to_ship",
		},
	},
	{
		Name: "Profile 6", // profile smart cuy //
		URLs: []string{
			"https://seller-id.tokopedia.com/order?order_status[]=1&selected_sort=11&tab=to_ship&urgency[]=201",
		},
	},
	{
		Name: "Profile 2", // profile mija //
		URLs: []string{
			"https://seller-id.tokopedia.com/order?order_status[]=1&selected_sort=11&tab=to_ship",
		},
	},
	{
		Name: "Profile 1", // profile shopee ajim //
		URLs: []string{
			"https://seller-id.tokopedia.com/order?order_status[]=1&selected_sort=11&tab=to_ship&urgency[]=201",
		},
	},
	{
		Name: "Profile 21", // profile shopee ajim //
		URLs: []string{
			" https://seller.blibli.com/order-list",
		},
	},
	{
		Name: "Profile 18", // profile shopee ajim //
		URLs: []string{
			" https://www.facebook.com/?locale=id_ID",
		},
	},
}
