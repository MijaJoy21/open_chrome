package main

import (
	"fmt"
	"os/exec"
	"time"
)

type ChromeProfile struct {
	Name string
	URLs []string
}

func main() {
	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`

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

	profiles := []ChromeProfile{
		{
			Name: "Profile 2", // profile xs kins  //
			URLs: []string{
				"https://web.whatsapp.com/",
				"https://seller.shopee.co.id/",
				"https://seller-id.tokopedia.com/product/manage?shop_region=ID",
			},
		},
		{
			Name: "Profile 4", // profile kins ungu //
			URLs: []string{
				"https://www.tokopedia.com/kins-official/product",
				"https://www.tokopedia.com/kinsofficialstore/product/",
			},
		},
		{
			Name: "Profile 5", // profile monkey tech //
			URLs: []string{
				"https://seller-id.tokopedia.com/product/manage?shop_region=ID",
				"https://seller-id.tokopedia.com/promotion/marketing-tools/management?tab=1&shop_region=ID",
			},
		},
		{
			Name: "Profile 19", // profile smart cuy //
			URLs: []string{
				"https://seller-id.tokopedia.com/product/manage?shop_region=ID",
				"https://seller-id.tokopedia.com/promotion/marketing-tools/management?tab=1&shop_region=ID",
			},
		},
		{
			Name: "Profile 10", // profile mija //
			URLs: []string{
				"https://seller-id.tokopedia.com/product/manage?shop_region=ID",
				"https://seller-id.tokopedia.com/promotion/marketing-tools/management?tab=1&shop_region=ID",
			},
		},
		{
			Name: "Profile 21", // profile shopee ajim //
			URLs: []string{
				"https://seller.shopee.co.id/portal/product/list/live/all?operationSortBy=modified_time&page=1",
				"https://seller.shopee.co.id/portal/marketing/list/discount?discountType=all&status=all&page=1",
			},
		},
	}

	for _, profile := range profiles {

		args := []string{
			"--new-window",
			"--profile-directory=" + profile.Name,
		}

		args = append(args, profile.URLs...)

		cmd := exec.Command(chromePath, args...)

		err := cmd.Start()
		if err != nil {
			fmt.Printf("Gagal membuka %s: %v\n", profile.Name, err)
		} else {
			fmt.Println("berhasil membuka", profile.Name)
		}

		time.Sleep(3 * time.Second)
	}

}
