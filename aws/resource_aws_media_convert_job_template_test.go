package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAwsMediaConvertJobTemplate_base(t *testing.T) {
	var jobTemplate mediaconvert.JobTemplate
	resourceName := "aws_media_convert_job-template.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-base")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertJobTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertJobTemplateConfig_Basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertJobTemplateExists(resourceName, &jobTemplate),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`jobtemplate/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func testAccMediaConvertJobTemplateConfig_Basic(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_job_template" "test" {
		name = %[1]q
		description = "test job template"
		acceleration_settings {
			mode = "DISABLED"
		}
		priority = 0
		settings {
			ad_avail_offset = 0
			input {
					audio_selector_group {
						name = "Audio Selector Group 1"
						audio_selector_names = ["Audio Selector 1"]
					}
					
					audio_selector {
							name = "Audio Selector 1"
							default_selection = "DEFAULT"
							offset = 0
							program_selection = 1
					}
					
					caption_selector {
							name = "Captions Selector 1"
							source_settings {
								embedded_source_settings {
									convert_608_to_708 = "DISABLED"
									source_608_channel_number = 1
									source_608_track_number = 1
									terminate_captions = "END_OF_INPUT"
								}
								source_type = "EMBEDDED"
							}
					}
										        
					deblock_filter = "DISABLED"
					denoise_filter = "DISABLED"
					filter_enable = "AUTO"
					filter_strength = 0
					psi_control = "USE_PSI"
					timecode_source = "EMBEDDED"
					video_selector {
						alpha_behavior = "DISCARD"
						color_space = "FOLLOW"
						rotate = "DEGREE_0"
					}
				}
		}
	}
	`, rName)
}

func testAccCheckAwsMediaConvertJobTemplateDestroy(s *terraform.State) error {
	//for _, rs := range s.RootModule().Resources {
	// 	if rs.Type != "aws_media_convert_preset" {
	// 		continue
	// 	}
	// 	conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
	// 	if err != nil {
	// 		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	// 	}

	// 	_, err = conn.GetPreset(&mediaconvert.GetPresetInput{
	// 		Name: aws.String(rs.Primary.ID),
	// 	})
	// 	if err != nil {
	// 		if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
	// 			continue
	// 		}
	// 		return err
	// 	}
	// }

	return nil
}

func testAccCheckAwsMediaConvertJobTemplateExists(n string, jobTemplate *mediaconvert.JobTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Preset id is set")
		}

		// conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		// if err != nil {
		// 	return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		// }

		// resp, err := conn.GetPreset(&mediaconvert.GetPresetInput{
		// 	Name: aws.String(rs.Primary.ID),
		// })
		// if err != nil {
		// 	return fmt.Errorf("Error getting preset: %s", err)
		// }

		// *preset = *resp.Preset
		return nil
	}
}
