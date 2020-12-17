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
  category = ""
  settings {
	video_description {
		scaling_behavior = "DEFAULT"
		timecode_insertion = "DISABLED"
		anti_alias = "ENABLED"
		sharpness = 50
		afd_signaling = "NONE"
		drop_frame_timecode = "ENABLED"
		respond_to_afd = "NONE"
		color_metadata = "INSERT"
		codec_settings {
			codec = "H_264"
			///https://docs.aws.amazon.com/sdk-for-go/api/service/mediaconvert/#H264Settings
			h264_settings {

			}
		}
	}
	audio_description {
		audio_type_control = "FOLLOW_INPUT"
		audio_source_name = "Audio Selector 1"
		language_code_control = "FOLLOW_INPUT"
		codec_settings {
			codec = "AAC"
			aac_settings {
				audio_description_broadcaster_mix = "NORMAL"
				bitrate = 96000
				rate_control_mode = "CBR"
				coding_mode = "CODING_MODE_2_0"
				raw_format = "NONE"
				sample_rate = 48000
				specification = "MPEG4"
			}
		}
	}

	container_settings {
		container = "MP4"
		mp4_settings {
			cslg_atom = "INCLUDE"
			ctts_version = 0
			free_space_box = "EXCLUDE"
			moov_placement = "PROGRESSIVE_DOWNLOAD"
		}
	}
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
