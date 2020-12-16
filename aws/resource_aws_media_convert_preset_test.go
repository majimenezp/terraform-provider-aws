package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAwsMediaConvertPreset_base(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test")
	rCategory := acctest.RandomWithPrefix("tf-acc-test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_Basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`preset/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "category", rCategory),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMediaConvertPresetConfig_Basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_media_convert_preset" "test" {
  name = %[1]q
  settings {

  }
}
`, rName)
}

// func testAccPreCheckAWSMediaConvert(t *testing.T) {
// 	_, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))

// 	if testAccPreCheckSkipError(err) {
// 		t.Skipf("skipping acceptance testing: %s", err)
// 	}

// 	if err != nil {
// 		t.Fatalf("unexpected PreCheck error: %s", err)
// 	}
// }

func testAccCheckAwsMediaConvertPresetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_media_convert_preset" {
			continue
		}
		conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		if err != nil {
			return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		}

		_, err = conn.GetPreset(&mediaconvert.GetPresetInput{
			Name: aws.String(rs.Primary.ID),
		})
		if err != nil {
			if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
				continue
			}
			return err
		}
	}

	return nil
}

func testAccCheckAwsMediaConvertPresetExists(n string, preset *mediaconvert.Preset) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Preset id is set")
		}

		conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		if err != nil {
			return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		}

		resp, err := conn.GetPreset(&mediaconvert.GetPresetInput{
			Name: aws.String(rs.Primary.ID),
		})
		if err != nil {
			return fmt.Errorf("Error getting preset: %s", err)
		}

		*preset = *resp.Preset
		return nil
	}
}
