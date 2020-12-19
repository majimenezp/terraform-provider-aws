package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
)

func resourceAwsMediaConvertPreset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsMediaConvertPresetCreate,
		Read:   resourceAwsMediaConvertPresetRead,
		Update: resourceAwsMediaConvertPresetUpdate,
		Delete: resourceAwsMediaConvertPresetDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"settings": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audio_description": {
							Type:     schema.TypeList,
							MinItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audio_source_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"audio_type": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"audio_type_control": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AudioTypeControlFollowInput,
											mediaconvert.AudioTypeControlUseConfigured,
										}, false),
									},
									"custom_language_code": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"language_code": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
									},
									"language_code_control": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AudioLanguageCodeControlFollowInput,
											mediaconvert.AudioLanguageCodeControlUseConfigured,
										}, false),
									},
									"stream_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"audio_channel_tagging_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channel_tag": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioChannelTagL,
														mediaconvert.AudioChannelTagR,
														mediaconvert.AudioChannelTagC,
														mediaconvert.AudioChannelTagLfe,
														mediaconvert.AudioChannelTagLs,
														mediaconvert.AudioChannelTagRs,
														mediaconvert.AudioChannelTagLc,
														mediaconvert.AudioChannelTagRc,
														mediaconvert.AudioChannelTagCs,
														mediaconvert.AudioChannelTagLsd,
														mediaconvert.AudioChannelTagRsd,
														mediaconvert.AudioChannelTagTcs,
														mediaconvert.AudioChannelTagVhl,
														mediaconvert.AudioChannelTagVhc,
														mediaconvert.AudioChannelTagVhr,
													}, false),
												},
											},
										},
									},
									"audio_normalization_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationAlgorithmItuBs17701,
														mediaconvert.AudioNormalizationAlgorithmItuBs17702,
														mediaconvert.AudioNormalizationAlgorithmItuBs17703,
														mediaconvert.AudioNormalizationAlgorithmItuBs17704,
													}, false),
												},
												"algorithm_control": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationAlgorithmControlCorrectAudio,
														mediaconvert.AudioNormalizationAlgorithmControlMeasureOnly,
													}, false),
												},
												"correction_gate_level": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"loudness_logging": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationLoudnessLoggingLog,
														mediaconvert.AudioNormalizationLoudnessLoggingDontLog,
													}, false),
												},
												"peak_calculation": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationPeakCalculationTruePeak,
														mediaconvert.AudioNormalizationPeakCalculationNone,
													}, false),
												},
												"target_lkfs": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
											},
										},
									},
									"codec_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"codec": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioCodecAac,
														mediaconvert.AudioCodecMp2,
														mediaconvert.AudioCodecMp3,
														mediaconvert.AudioCodecWav,
														mediaconvert.AudioCodecAiff,
														mediaconvert.AudioCodecAc3,
														mediaconvert.AudioCodecEac3,
														mediaconvert.AudioCodecEac3Atmos,
														mediaconvert.AudioCodecVorbis,
														mediaconvert.AudioCodecOpus,
														mediaconvert.AudioCodecPassthrough,
													}, false),
												},
												"aac_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audio_description_broadcaster_mix": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacAudioDescriptionBroadcasterMixBroadcasterMixedAd,
																	mediaconvert.AacAudioDescriptionBroadcasterMixNormal,
																}, false),
																Default: nil,
															},
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(6000),
																Default:      6000,
															},
															"codec_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacCodecProfileLc,
																	mediaconvert.AacCodecProfileHev1,
																	mediaconvert.AacCodecProfileHev2,
																}, false),
																Default: nil,
															},
															"coding_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacCodingModeAdReceiverMix,
																	mediaconvert.AacCodingModeCodingMode10,
																	mediaconvert.AacCodingModeCodingMode11,
																	mediaconvert.AacCodingModeCodingMode20,
																	mediaconvert.AacCodingModeCodingMode51,
																}, false),
																Default: nil,
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacRateControlModeCbr,
																	mediaconvert.AacRateControlModeVbr,
																}, false),
																Default: nil,
															},
															"raw_format": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacRawFormatLatmLoas,
																	mediaconvert.AacRawFormatNone,
																}, false),
																Default: nil,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(8000),
																Default:      8000,
															},
															"specification": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacSpecificationMpeg2,
																	mediaconvert.AacSpecificationMpeg4,
																}, false),
																Default: nil,
															},
															"vbr_quality": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AacVbrQualityLow,
																	mediaconvert.AacVbrQualityMediumLow,
																	mediaconvert.AacVbrQualityMediumHigh,
																	mediaconvert.AacVbrQualityHigh,
																}, false),
																Default: nil,
															},
														},
													},
												},
												"ac3_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(64000),
																Default:      64000,
															},
															"bitstream_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Ac3BitstreamModeCompleteMain,
																	mediaconvert.Ac3BitstreamModeCommentary,
																	mediaconvert.Ac3BitstreamModeDialogue,
																	mediaconvert.Ac3BitstreamModeEmergency,
																	mediaconvert.Ac3BitstreamModeHearingImpaired,
																	mediaconvert.Ac3BitstreamModeMusicAndEffects,
																	mediaconvert.Ac3BitstreamModeVisuallyImpaired,
																	mediaconvert.Ac3BitstreamModeVoiceOver,
																}, false),
																Default: nil,
															},
															"coding_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Ac3CodingModeCodingMode10,
																	mediaconvert.Ac3CodingModeCodingMode11,
																	mediaconvert.Ac3CodingModeCodingMode20,
																	mediaconvert.Ac3CodingModeCodingMode32Lfe,
																}, false),
																Default: nil,
															},
															"dialnorm": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"dynamic_range_compression_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Ac3DynamicRangeCompressionProfileFilmStandard,
																	mediaconvert.Ac3DynamicRangeCompressionProfileNone,
																}, false),
																Default: nil,
															},
															"lfe_filter": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Ac3LfeFilterEnabled,
																	mediaconvert.Ac3LfeFilterDisabled,
																}, false),
																Default: nil,
															},
															"metadata_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Ac3MetadataControlFollowInput,
																	mediaconvert.Ac3MetadataControlUseConfigured,
																}, false),
																Default: nil,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(48000),
																Default:      48000,
															},
														},
													},
												},
												"aiff_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitdepth": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(16),
																Default:      16,
															},
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(8000),
																Default:      8000,
															},
														},
													},
												},
												"eac3_atmos_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(384000),
																Default:      384000,
															},
															"bitstream_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosBitstreamModeCompleteMain,
																}, false),
															},
															"coding_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosCodingModeCodingMode916,
																}, false),
															},
															"dialogue_intelligence": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosDialogueIntelligenceEnabled,
																	mediaconvert.Eac3AtmosDialogueIntelligenceDisabled,
																}, false),
															},
															"dynamic_range_compression_line": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineNone,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineFilmStandard,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineFilmLight,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineMusicStandard,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineMusicLight,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionLineSpeech,
																}, false),
															},
															"dynamic_range_compression_rf": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfNone,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfFilmStandard,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfFilmLight,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfMusicStandard,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfMusicLight,
																	mediaconvert.Eac3AtmosDynamicRangeCompressionRfSpeech,
																}, false),
															},
															"lo_ro_center_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lo_ro_surround_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lt_rt_center_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lt_rt_surround_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"metering_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosMeteringModeLeqA,
																	mediaconvert.Eac3AtmosMeteringModeItuBs17701,
																	mediaconvert.Eac3AtmosMeteringModeItuBs17702,
																	mediaconvert.Eac3AtmosMeteringModeItuBs17703,
																	mediaconvert.Eac3AtmosMeteringModeItuBs17704,
																}, false),
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(48000),
																Default:      48000,
															},
															"speech_threshold": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"stereo_downmix": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosStereoDownmixNotIndicated,
																	mediaconvert.Eac3AtmosStereoDownmixStereo,
																	mediaconvert.Eac3AtmosStereoDownmixSurround,
																	mediaconvert.Eac3AtmosStereoDownmixDpl2,
																}, false),
															},
															"surround_ex_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AtmosSurroundExModeNotIndicated,
																	mediaconvert.Eac3AtmosSurroundExModeEnabled,
																	mediaconvert.Eac3AtmosSurroundExModeDisabled,
																}, false),
															},
														},
													},
												},
												"eac3_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attenuation_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3AttenuationControlAttenuate3Db,
																	mediaconvert.Eac3AttenuationControlNone,
																}, false),
															},
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(64000),
																Default:      64000,
															},
															"bitstream_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3BitstreamModeCompleteMain,
																	mediaconvert.Eac3BitstreamModeCommentary,
																	mediaconvert.Eac3BitstreamModeEmergency,
																	mediaconvert.Eac3BitstreamModeHearingImpaired,
																	mediaconvert.Eac3BitstreamModeVisuallyImpaired,
																}, false),
															},
															"coding_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3CodingModeCodingMode10,
																	mediaconvert.Eac3CodingModeCodingMode20,
																	mediaconvert.Eac3CodingModeCodingMode32,
																}, false),
															},
															"dc_filter": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3DcFilterEnabled,
																	mediaconvert.Eac3DcFilterDisabled,
																}, false),
															},
															"dialnorm": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"dynamic_range_compression_line": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3DynamicRangeCompressionLineNone,
																	mediaconvert.Eac3DynamicRangeCompressionLineFilmStandard,
																	mediaconvert.Eac3DynamicRangeCompressionLineFilmLight,
																	mediaconvert.Eac3DynamicRangeCompressionLineMusicStandard,
																	mediaconvert.Eac3DynamicRangeCompressionLineMusicLight,
																	mediaconvert.Eac3DynamicRangeCompressionLineSpeech,
																}, false),
															},
															"dynamic_range_compression_rf": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3DynamicRangeCompressionRfNone,
																	mediaconvert.Eac3DynamicRangeCompressionRfFilmStandard,
																	mediaconvert.Eac3DynamicRangeCompressionRfFilmLight,
																	mediaconvert.Eac3DynamicRangeCompressionRfMusicStandard,
																	mediaconvert.Eac3DynamicRangeCompressionRfMusicLight,
																	mediaconvert.Eac3DynamicRangeCompressionRfSpeech,
																}, false),
															},
															"lfe_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3LfeControlLfe,
																	mediaconvert.Eac3LfeControlNoLfe,
																}, false),
															},
															"lfe_filter": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3LfeFilterEnabled,
																	mediaconvert.Eac3LfeFilterDisabled,
																}, false),
															},
															"lo_ro_center_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lo_ro_surround_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lt_rt_center_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"lt_rt_surround_mix_level": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"metadata_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3MetadataControlFollowInput,
																	mediaconvert.Eac3MetadataControlUseConfigured,
																}, false),
															},
															"passthrough_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3PassthroughControlWhenPossible,
																	mediaconvert.Eac3PassthroughControlNoPassthrough,
																}, false),
															},
															"phase_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3PhaseControlShift90Degrees,
																	mediaconvert.Eac3PhaseControlNoShift,
																}, false),
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(48000),
																Default:      48000,
															},
															"stereo_downmix": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3StereoDownmixNotIndicated,
																	mediaconvert.Eac3StereoDownmixLoRo,
																	mediaconvert.Eac3StereoDownmixLtRt,
																	mediaconvert.Eac3StereoDownmixDpl2,
																}, false),
															},
															"surround_ex_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3SurroundExModeNotIndicated,
																	mediaconvert.Eac3SurroundExModeEnabled,
																	mediaconvert.Eac3SurroundExModeDisabled,
																}, false),
															},
															"surround_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Eac3SurroundModeNotIndicated,
																	mediaconvert.Eac3SurroundModeEnabled,
																	mediaconvert.Eac3SurroundModeDisabled,
																}, false),
															},
														},
													},
												},
												"mp2_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32000),
																Default:      32000,
															},
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32000),
																Default:      32000,
															},
														},
													},
												},
												"mp3_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(16000),
																Default:      16000,
															},
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mp3RateControlModeCbr,
																	mediaconvert.Mp3RateControlModeVbr,
																}, false),
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(22050),
																Default:      22050,
															},
															"vbr_quality": {
																Type:     schema.TypeInt,
																Optional: true,
																Default:  0,
															},
														},
													},
												},
												"opus_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32000),
																Default:      32000,
															},
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(16000),
																Default:      16000,
															},
														},
													},
												},
												"vorbis_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"sample_rate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(22050),
																Default:      22050,
															},
															"vbr_quality": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"wav_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitdepth": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(16),
																Default:      16,
															},
															"channels": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
																Default:      1,
															},
															"format": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.WavFormatRiff,
																	mediaconvert.WavFormatRf64,
																}, false),
															},
															"sample_rate": {
																Optional:     true,
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(8000),
																Default:      8000,
															},
														},
													},
												},
											},
										},
									},
									"remix_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channel_mapping": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"output_channels": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"input_channels": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeInt},
																		},
																	},
																},
															},
														},
													},
												},
												"channels_in": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
												"channels_out": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
											},
										},
									},
								},
							},
						},
						"caption_description": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_language_code": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"destination_settings": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"burnin_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleAlignmentCentered,
																	mediaconvert.BurninSubtitleAlignmentLeft,
																}, false),
															},
															"background_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleBackgroundColorNone,
																	mediaconvert.BurninSubtitleBackgroundColorBlack,
																	mediaconvert.BurninSubtitleBackgroundColorWhite,
																}, false),
															},
															"background_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"font_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleFontColorWhite,
																	mediaconvert.BurninSubtitleFontColorBlack,
																	mediaconvert.BurninSubtitleFontColorYellow,
																	mediaconvert.BurninSubtitleFontColorRed,
																	mediaconvert.BurninSubtitleFontColorGreen,
																	mediaconvert.BurninSubtitleFontColorBlue,
																}, false),
															},
															"font_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"font_resolution": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(96),
															},
															"font_script": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.FontScriptAutomatic,
																	mediaconvert.FontScriptHans,
																	mediaconvert.FontScriptHant,
																}, false),
															},
															"font_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"outline_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleOutlineColorBlack,
																	mediaconvert.BurninSubtitleOutlineColorWhite,
																	mediaconvert.BurninSubtitleOutlineColorYellow,
																	mediaconvert.BurninSubtitleOutlineColorRed,
																	mediaconvert.BurninSubtitleOutlineColorGreen,
																	mediaconvert.BurninSubtitleOutlineColorBlue,
																}, false),
															},
															"outline_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleShadowColorNone,
																	mediaconvert.BurninSubtitleShadowColorBlack,
																	mediaconvert.BurninSubtitleShadowColorWhite,
																}, false),
															},
															"shadow_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_x_offset": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_y_offset": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"teletext_spacing": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleTeletextSpacingFixedGrid,
																	mediaconvert.BurninSubtitleTeletextSpacingProportional,
																}, false),
															},
															"x_position": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"y_position": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"destination_type": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CaptionDestinationTypeBurnIn,
														mediaconvert.CaptionDestinationTypeDvbSub,
														mediaconvert.CaptionDestinationTypeEmbedded,
														mediaconvert.CaptionDestinationTypeEmbeddedPlusScte20,
														mediaconvert.CaptionDestinationTypeImsc,
														mediaconvert.CaptionDestinationTypeScte20PlusEmbedded,
														mediaconvert.CaptionDestinationTypeScc,
														mediaconvert.CaptionDestinationTypeSrt,
														mediaconvert.CaptionDestinationTypeSmi,
														mediaconvert.CaptionDestinationTypeTeletext,
														mediaconvert.CaptionDestinationTypeTtml,
														mediaconvert.CaptionDestinationTypeWebvtt,
													}, false),
												},
												"dvb_sub_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleAlignmentCentered,
																	mediaconvert.DvbSubtitleAlignmentLeft,
																}, false),
															},
															"background_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleBackgroundColorNone,
																	mediaconvert.DvbSubtitleBackgroundColorBlack,
																	mediaconvert.DvbSubtitleBackgroundColorWhite,
																}, false),
															},
															"background_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"font_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleFontColorWhite,
																	mediaconvert.DvbSubtitleFontColorBlack,
																	mediaconvert.DvbSubtitleFontColorYellow,
																	mediaconvert.DvbSubtitleFontColorRed,
																	mediaconvert.DvbSubtitleFontColorGreen,
																	mediaconvert.DvbSubtitleFontColorBlue,
																}, false),
															},
															"font_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"font_resolution": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(96),
															},
															"font_script": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.FontScriptAutomatic,
																	mediaconvert.FontScriptHans,
																	mediaconvert.FontScriptHant,
																}, false),
															},
															"font_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"outline_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleOutlineColorBlack,
																	mediaconvert.DvbSubtitleOutlineColorWhite,
																	mediaconvert.DvbSubtitleOutlineColorYellow,
																	mediaconvert.DvbSubtitleOutlineColorRed,
																	mediaconvert.DvbSubtitleOutlineColorGreen,
																	mediaconvert.DvbSubtitleOutlineColorBlue,
																}, false),
															},
															"outline_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_color": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleShadowColorNone,
																	mediaconvert.DvbSubtitleShadowColorBlack,
																	mediaconvert.DvbSubtitleShadowColorWhite,
																}, false),
															},
															"shadow_opacity": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_x_offset": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"shadow_y_offset": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"subtitling_type": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitlingTypeHearingImpaired,
																	mediaconvert.DvbSubtitlingTypeStandard,
																}, false),
															},
															"teletext_spacing": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleTeletextSpacingFixedGrid,
																	mediaconvert.DvbSubtitleTeletextSpacingProportional,
																}, false),
															},
															"x_position": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"y_position": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"embedded_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination_608_channel_number": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"destination_708_service_number": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
														},
													},
												},
												"imsc_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_passthrough": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ImscStylePassthroughEnabled,
																	mediaconvert.ImscStylePassthroughDisabled,
																}, false),
															},
														},
													},
												},
												"scc_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"framerate": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.SccDestinationFramerateFramerate2397,
																	mediaconvert.SccDestinationFramerateFramerate24,
																	mediaconvert.SccDestinationFramerateFramerate25,
																	mediaconvert.SccDestinationFramerateFramerate2997Dropframe,
																	mediaconvert.SccDestinationFramerateFramerate2997NonDropframe,
																}, false),
															},
														},
													},
												},
												"teletext_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"page_number": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(3, 256),
															},
															"page_types": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Set:      schema.HashString,
															},
														},
													},
												},
												"ttml_destination_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_passthrough": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.TtmlStylePassthroughEnabled,
																	mediaconvert.TtmlStylePassthroughDisabled,
																}, false),
															},
														},
													},
												},
											},
										},
									},
									"language_code": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
									},
									"language_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"container_settings": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmfc_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcAudioDurationDefaultCodecDuration,
														mediaconvert.CmfcAudioDurationMatchVideoDuration,
													}, false),
												},
												"scte35_esam": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcScte35EsamInsert,
														mediaconvert.CmfcScte35EsamNone,
													}, false),
												},
												"scte35_source": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcScte35SourcePassthrough,
														mediaconvert.CmfcScte35SourceNone,
													}, false),
												},
											},
										},
									},
									"container": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.ContainerTypeF4v,
											mediaconvert.ContainerTypeIsmv,
											mediaconvert.ContainerTypeM2ts,
											mediaconvert.ContainerTypeM3u8,
											mediaconvert.ContainerTypeCmfc,
											mediaconvert.ContainerTypeMov,
											mediaconvert.ContainerTypeMp4,
											mediaconvert.ContainerTypeMpd,
											mediaconvert.ContainerTypeMxf,
											mediaconvert.ContainerTypeWebm,
											mediaconvert.ContainerTypeRaw,
										}, false),
									},
									"f4v_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"moov_placement": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.F4vMoovPlacementProgressiveDownload,
														mediaconvert.F4vMoovPlacementNormal,
													}, false),
												},
											},
										},
									},
									"m2ts_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsAudioDurationDefaultCodecDuration,
														mediaconvert.M2tsAudioDurationMatchVideoDuration,
													}, false),
												},
												"audio_frames_per_pes": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"audio_pids": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeInt},
													Set:      schema.HashString,
												},
												"bitrate": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"buffer_model": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsBufferModelMultiplex,
														mediaconvert.M2tsBufferModelNone,
													}, false),
												},
												"dvb_nit_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_id": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"network_name": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
															"nit_interval": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(25),
															},
														},
													},
												},
												"dvb_sdt_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"output_sdt": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.OutputSdtSdtFollow,
																	mediaconvert.OutputSdtSdtFollowIfPresent,
																	mediaconvert.OutputSdtSdtManual,
																	mediaconvert.OutputSdtSdtNone,
																}, false),
															},
															"sdt_interval": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(25),
															},
															"service_name": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
															"service_provider_name": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
														},
													},
												},
												"dvb_sub_pids": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeInt},
													Set:      schema.HashString,
												},
												"dvb_tdt_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tdt_interval": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
														},
													},
												},
												"dvb_teletext_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      499,
												},
												"ebp_audio_interval": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEbpAudioIntervalVideoAndFixedIntervals,
														mediaconvert.M2tsEbpAudioIntervalVideoInterval,
													}, false),
												},
												"ebp_placement": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEbpPlacementVideoAndAudioPids,
														mediaconvert.M2tsEbpPlacementVideoPid,
													}, false),
												},
												"es_rate_in_pes": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEsRateInPesInclude,
														mediaconvert.M2tsEsRateInPesExclude,
													}, false),
												},
												"force_ts_video_ebp_order": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsForceTsVideoEbpOrderForce,
														mediaconvert.M2tsForceTsVideoEbpOrderDefault,
													}, false),
												},
												"fragment_time": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"max_pcr_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"min_pcr_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"nielsen_id3": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsNielsenId3Insert,
														mediaconvert.M2tsNielsenId3None,
													}, false),
												},
												"null_packet_bitrate": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"pat_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"pcr_control": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsPcrControlPcrEveryPesPacket,
														mediaconvert.M2tsPcrControlConfiguredPcrPeriod,
													}, false),
												},
												"pcr_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"pmt_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"pmt_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      48,
												},
												"private_metadata_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      503,
												},
												"program_number": {
													Type:     schema.TypeInt,
													Optional: true,
													Default:  1,
												},
												"rate_mode": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsRateModeVbr,
														mediaconvert.M2tsRateModeCbr,
													}, false),
												},
												"scte_35_esam": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"scte_35_esam_pid": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32),
															},
														},
													},
												},
												"scte_35_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"scte_35_source": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsScte35SourcePassthrough,
														mediaconvert.M2tsScte35SourceNone,
													}, false),
												},
												"segmentation_markers": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsSegmentationMarkersNone,
														mediaconvert.M2tsSegmentationMarkersRaiSegstart,
														mediaconvert.M2tsSegmentationMarkersRaiAdapt,
														mediaconvert.M2tsSegmentationMarkersPsiSegstart,
														mediaconvert.M2tsSegmentationMarkersEbp,
														mediaconvert.M2tsSegmentationMarkersEbpLegacy,
													}, false),
												},
												"segmentation_style": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsSegmentationStyleMaintainCadence,
														mediaconvert.M2tsSegmentationStyleResetCadence,
													}, false),
												},
												"segmentation_time": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"timed_metadata_pid": {
													Type:     schema.TypeInt,
													Optional: true,
													Default:  32,
												},
												"transport_stream_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"video_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
											},
										},
									},
									"m3u8_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8AudioDurationDefaultCodecDuration,
														mediaconvert.M3u8AudioDurationMatchVideoDuration,
													}, false)},
												"audio_frames_per_pes": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"audio_pids": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeInt},
													Set:      schema.HashString,
												},
												"nielsen_id3": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8NielsenId3Insert,
														mediaconvert.M3u8NielsenId3None,
													}, false),
												},
												"pat_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"pcr_control": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8PcrControlPcrEveryPesPacket,
														mediaconvert.M3u8PcrControlConfiguredPcrPeriod,
													}, false),
												},
												"pcr_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"pmt_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"pmt_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"private_metadata_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"program_number": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"scte_35_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"scte_35_source": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8Scte35SourcePassthrough,
														mediaconvert.M3u8Scte35SourceNone,
													}, false),
												},
												"timed_metadata": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.TimedMetadataPassthrough,
														mediaconvert.TimedMetadataNone,
													}, false),
												},
												"timed_metadata_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"transport_stream_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"video_pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(32),
												},
											},
										},
									},
									"mov_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"clap_atom": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovClapAtomInclude,
														mediaconvert.MovClapAtomExclude,
													}, false),
												},
												"cslg_atom": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovCslgAtomInclude,
														mediaconvert.MovCslgAtomExclude,
													}, false),
												},
												"mpeg2_fourcc_control": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovMpeg2FourCCControlXdcam,
														mediaconvert.MovMpeg2FourCCControlMpeg,
													}, false),
												},
												"padding_control": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovPaddingControlOmneon,
														mediaconvert.MovPaddingControlNone,
													}, false),
												},
												"reference": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovReferenceSelfContained,
														mediaconvert.MovReferenceExternal,
													}, false),
												},
											},
										},
									},
									"mp4_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcAudioDurationDefaultCodecDuration,
														mediaconvert.CmfcAudioDurationMatchVideoDuration,
													}, false),
												},
												"cslg_atom": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4CslgAtomInclude,
														mediaconvert.Mp4CslgAtomExclude,
													}, false),
												},
												"ctts_version": {
													Type:     schema.TypeInt,
													Optional: true,
													Default:  0,
												},
												"free_space_box": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4FreeSpaceBoxInclude,
														mediaconvert.Mp4FreeSpaceBoxExclude,
													}, false),
												},
												"moov_placement": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4MoovPlacementProgressiveDownload,
														mediaconvert.Mp4MoovPlacementNormal,
													}, false),
												},
												"mp4_major_brand": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"mpd_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility_caption_hints": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdAccessibilityCaptionHintsInclude,
														mediaconvert.MpdAccessibilityCaptionHintsExclude,
													}, false),
												},
												"audio_duration": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdAudioDurationDefaultCodecDuration,
														mediaconvert.MpdAudioDurationMatchVideoDuration,
													}, false)},
												"caption_container_type": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdCaptionContainerTypeRaw,
														mediaconvert.MpdCaptionContainerTypeFragmentedMp4,
													}, false),
												},
												"scte_35_esam": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdScte35EsamInsert,
														mediaconvert.MpdScte35EsamNone,
													}, false),
												},
												"scte_35_source": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdScte35SourcePassthrough,
														mediaconvert.MpdScte35SourceNone,
													}, false),
												},
											},
										},
									},
									"mxf_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"afd_signaling": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MxfAfdSignalingNoCopy,
														mediaconvert.MxfAfdSignalingCopyFromVideo,
													}, false),
												},
												"profile": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MxfProfileD10,
														mediaconvert.MxfProfileXdcam,
														mediaconvert.MxfProfileOp1a,
													}, false),
												},
											},
										},
									},
								},
							},
						},
						"video_description": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"afd_signaling": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AfdSignalingNone,
											mediaconvert.AfdSignalingAuto,
											mediaconvert.AfdSignalingFixed,
										}, false),
									},
									"anti_alias": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AntiAliasDisabled,
											mediaconvert.AntiAliasEnabled,
										}, false),
									},
									"codec_settings": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"av1_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Av1AdaptiveQuantizationOff,
																	mediaconvert.Av1AdaptiveQuantizationLow,
																	mediaconvert.Av1AdaptiveQuantizationMedium,
																	mediaconvert.Av1AdaptiveQuantizationHigh,
																	mediaconvert.Av1AdaptiveQuantizationHigher,
																	mediaconvert.Av1AdaptiveQuantizationMax,
																}, false),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Av1FramerateControlInitializeFromSource,
																	mediaconvert.Av1FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Av1FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.Av1FramerateConversionAlgorithmInterpolate,
																	mediaconvert.Av1FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"gop_size": {
																Type:         schema.TypeFloat,
																Optional:     true,
																ValidateFunc: validation.FloatAtLeast(0),
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"number_b_frames_between_reference_frames": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntBetween(7, 15),
															},
															"qvbr_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"qvbr_quality_level": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																		"qvbr_quality_level_fine_tune": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																			Default:  0,
																		},
																	},
																},
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Av1RateControlModeQvbr,
																}, false),
															},
															"slices": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"spatial_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Av1SpatialAdaptiveQuantizationDisabled,
																	mediaconvert.Av1SpatialAdaptiveQuantizationEnabled,
																}, false),
																Default: mediaconvert.Av1SpatialAdaptiveQuantizationEnabled,
															},
														},
													},
												},
												"avc_intra_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"avc_intra_class": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraClassClass50,
																	mediaconvert.AvcIntraClassClass100,
																	mediaconvert.AvcIntraClassClass200,
																}, false),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraFramerateControlInitializeFromSource,
																	mediaconvert.AvcIntraFramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraFramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.AvcIntraFramerateConversionAlgorithmInterpolate,
																	mediaconvert.AvcIntraFramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(24),
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraInterlaceModeProgressive,
																	mediaconvert.AvcIntraInterlaceModeTopField,
																	mediaconvert.AvcIntraInterlaceModeBottomField,
																	mediaconvert.AvcIntraInterlaceModeFollowTopField,
																	mediaconvert.AvcIntraInterlaceModeFollowBottomField,
																}, false),
															},
															"slow_pal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraSlowPalDisabled,
																	mediaconvert.AvcIntraSlowPalEnabled,
																}, false),
																Default: mediaconvert.AvcIntraSlowPalDisabled,
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.AvcIntraTelecineNone,
																	mediaconvert.AvcIntraTelecineHard,
																}, false),
																Default: mediaconvert.AvcIntraTelecineNone,
															},
														},
													},
												},
												"codec": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.VideoCodecAv1,
														mediaconvert.VideoCodecAvcIntra,
														mediaconvert.VideoCodecFrameCapture,
														mediaconvert.VideoCodecH264,
														mediaconvert.VideoCodecH265,
														mediaconvert.VideoCodecMpeg2,
														mediaconvert.VideoCodecProres,
														mediaconvert.VideoCodecVc3,
														mediaconvert.VideoCodecVp8,
														mediaconvert.VideoCodecVp9,
													}, false),
												},
												"frame_capture_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"max_captures": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
														},
													},
												},
												"h264_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264AdaptiveQuantizationOff,
																	mediaconvert.H264AdaptiveQuantizationAuto,
																	mediaconvert.H264AdaptiveQuantizationLow,
																	mediaconvert.H264AdaptiveQuantizationMedium,
																	mediaconvert.H264AdaptiveQuantizationHigh,
																	mediaconvert.H264AdaptiveQuantizationHigher,
																	mediaconvert.H264AdaptiveQuantizationMax,
																}, false),
															},
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
																Default:      1000,
															},
															"codec_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264CodecLevelAuto,
																	mediaconvert.H264CodecLevelLevel1,
																	mediaconvert.H264CodecLevelLevel11,
																	mediaconvert.H264CodecLevelLevel12,
																	mediaconvert.H264CodecLevelLevel13,
																	mediaconvert.H264CodecLevelLevel2,
																	mediaconvert.H264CodecLevelLevel21,
																	mediaconvert.H264CodecLevelLevel22,
																	mediaconvert.H264CodecLevelLevel3,
																	mediaconvert.H264CodecLevelLevel31,
																	mediaconvert.H264CodecLevelLevel32,
																	mediaconvert.H264CodecLevelLevel4,
																	mediaconvert.H264CodecLevelLevel41,
																	mediaconvert.H264CodecLevelLevel42,
																	mediaconvert.H264CodecLevelLevel5,
																	mediaconvert.H264CodecLevelLevel51,
																	mediaconvert.H264CodecLevelLevel52,
																}, false),
															},
															"codec_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264CodecProfileBaseline,
																	mediaconvert.H264CodecProfileHigh,
																	mediaconvert.H264CodecProfileHigh10bit,
																	mediaconvert.H264CodecProfileHigh422,
																	mediaconvert.H264CodecProfileHigh42210bit,
																	mediaconvert.H264CodecProfileMain,
																}, false),
															},
															"dynamic_sub_gop": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264DynamicSubGopAdaptive,
																	mediaconvert.H264DynamicSubGopStatic,
																}, false),
															},
															"entropy_encoding": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264EntropyEncodingCabac,
																	mediaconvert.H264EntropyEncodingCavlc,
																}, false),
															},
															"field_encoding": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264FieldEncodingPaff,
																	mediaconvert.H264FieldEncodingForceField,
																}, false),
															},
															"flicker_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264FlickerAdaptiveQuantizationDisabled,
																	mediaconvert.H264FlickerAdaptiveQuantizationEnabled,
																}, false),
																Default: mediaconvert.H264FlickerAdaptiveQuantizationEnabled,
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264FramerateControlInitializeFromSource,
																	mediaconvert.H264FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.H264FramerateConversionAlgorithmInterpolate,
																	mediaconvert.H264FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"gop_b_reference": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264GopBReferenceDisabled,
																	mediaconvert.H264GopBReferenceEnabled,
																}, false),
															},
															"gop_closed_cadence": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"gop_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"gop_size_units": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264GopSizeUnitsFrames,
																	mediaconvert.H264GopSizeUnitsSeconds,
																}, false),
															},
															"hrd_buffer_initial_fill_percentage": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"hrd_buffer_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264InterlaceModeProgressive,
																	mediaconvert.H264InterlaceModeTopField,
																	mediaconvert.H264InterlaceModeBottomField,
																	mediaconvert.H264InterlaceModeFollowTopField,
																	mediaconvert.H264InterlaceModeFollowBottomField,
																}, false),
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"min_i_interval": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"number_b_frames_between_reference_frames": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"number_reference_frames": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264ParControlInitializeFromSource,
																	mediaconvert.H264ParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality_tuning_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264QualityTuningLevelSinglePass,
																	mediaconvert.H264QualityTuningLevelSinglePassHq,
																	mediaconvert.H264QualityTuningLevelMultiPassHq,
																}, false),
															},
															"qvbr_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_average_bitrate": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1000),
																		},
																		"qvbr_quality_level": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntBetween(1, 10),
																		},
																		"qvbr_quality_level_fine_tune": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																	},
																},
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264RateControlModeVbr,
																	mediaconvert.H264RateControlModeCbr,
																	mediaconvert.H264RateControlModeQvbr,
																}, false),
															},
															"repeat_pps": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264RepeatPpsDisabled,
																	mediaconvert.H264RepeatPpsEnabled,
																}, false),
															},
															"scene_change_detect": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264SceneChangeDetectDisabled,
																	mediaconvert.H264SceneChangeDetectEnabled,
																	mediaconvert.H264SceneChangeDetectTransitionDetection,
																}, false),
															},
															"slices": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"slow_pal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264SlowPalDisabled,
																	mediaconvert.H264SlowPalEnabled,
																}, false),
																Default: mediaconvert.H264SlowPalDisabled,
															},
															"softness": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntBetween(0, 128),
															},
															"spatial_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264SpatialAdaptiveQuantizationDisabled,
																	mediaconvert.H264SpatialAdaptiveQuantizationEnabled,
																}, false),
															},
															"syntax": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264SyntaxDefault,
																	mediaconvert.H264SyntaxRp2027,
																}, false),
																Default: mediaconvert.H264SyntaxDefault,
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264TelecineNone,
																	mediaconvert.H264TelecineSoft,
																	mediaconvert.H264TelecineHard,
																}, false),
																Default: mediaconvert.Mpeg2TelecineNone,
															},
															"temporal_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264TemporalAdaptiveQuantizationDisabled,
																	mediaconvert.H264TemporalAdaptiveQuantizationEnabled,
																}, false),
																Default: mediaconvert.H264TemporalAdaptiveQuantizationEnabled,
															},
															"unregistered_sei_timecode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H264UnregisteredSeiTimecodeDisabled,
																	mediaconvert.H264UnregisteredSeiTimecodeEnabled,
																}, false),
															},
														},
													},
												},
												"h265_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265AdaptiveQuantizationOff,
																	mediaconvert.H265AdaptiveQuantizationLow,
																	mediaconvert.H265AdaptiveQuantizationMedium,
																	mediaconvert.H265AdaptiveQuantizationHigh,
																	mediaconvert.H265AdaptiveQuantizationHigher,
																	mediaconvert.H265AdaptiveQuantizationMax,
																}, false),
															},
															"alternate_transfer_function_sei": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265AlternateTransferFunctionSeiDisabled,
																	mediaconvert.H265AlternateTransferFunctionSeiEnabled,
																}, false),
															},
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"codec_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265CodecLevelAuto,
																	mediaconvert.H265CodecLevelLevel1,
																	mediaconvert.H265CodecLevelLevel2,
																	mediaconvert.H265CodecLevelLevel21,
																	mediaconvert.H265CodecLevelLevel3,
																	mediaconvert.H265CodecLevelLevel31,
																	mediaconvert.H265CodecLevelLevel4,
																	mediaconvert.H265CodecLevelLevel41,
																	mediaconvert.H265CodecLevelLevel5,
																	mediaconvert.H265CodecLevelLevel51,
																	mediaconvert.H265CodecLevelLevel52,
																	mediaconvert.H265CodecLevelLevel6,
																	mediaconvert.H265CodecLevelLevel61,
																	mediaconvert.H265CodecLevelLevel62,
																}, false),
															},
															"codec_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265CodecProfileMainMain,
																	mediaconvert.H265CodecProfileMainHigh,
																	mediaconvert.H265CodecProfileMain10Main,
																	mediaconvert.H265CodecProfileMain10High,
																	mediaconvert.H265CodecProfileMain4228bitMain,
																	mediaconvert.H265CodecProfileMain4228bitHigh,
																	mediaconvert.H265CodecProfileMain42210bitMain,
																	mediaconvert.H265CodecProfileMain42210bitHigh,
																}, false),
															},
															"dynamic_sub_gop": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265DynamicSubGopAdaptive,
																	mediaconvert.H265DynamicSubGopStatic,
																}, false),
															},
															"flicker_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265FlickerAdaptiveQuantizationDisabled,
																	mediaconvert.H265FlickerAdaptiveQuantizationEnabled,
																}, false),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265FramerateControlInitializeFromSource,
																	mediaconvert.H265FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.H265FramerateConversionAlgorithmInterpolate,
																	mediaconvert.H265FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"gop_b_reference": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265GopBReferenceDisabled,
																	mediaconvert.H265GopBReferenceEnabled,
																}, false),
															},
															"gop_closed_cadence": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"gop_size": {
																Type:     schema.TypeFloat,
																Optional: true,
															},
															"gop_size_units": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265GopSizeUnitsFrames,
																	mediaconvert.H265GopSizeUnitsSeconds,
																}, false),
															},
															"hrd_buffer_initial_fill_percentage": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"hrd_buffer_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265InterlaceModeProgressive,
																	mediaconvert.H265InterlaceModeTopField,
																	mediaconvert.H265InterlaceModeBottomField,
																	mediaconvert.H265InterlaceModeFollowTopField,
																	mediaconvert.H265InterlaceModeFollowBottomField,
																}, false),
																Default: mediaconvert.H265InterlaceModeProgressive,
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"min_i_interval": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"number_b_frames_between_reference_frames": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"number_reference_frames": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265ParControlInitializeFromSource,
																	mediaconvert.H265ParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality_tuning_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265QualityTuningLevelSinglePass,
																	mediaconvert.H265QualityTuningLevelSinglePassHq,
																	mediaconvert.H265QualityTuningLevelMultiPassHq,
																}, false),
															},
															"qvbr_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_average_bitrate": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1000),
																		},
																		"qvbr_quality_level": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntBetween(1, 10),
																		},
																		"qvbr_quality_level_fine_tune": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																	},
																},
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265RateControlModeVbr,
																	mediaconvert.H265RateControlModeCbr,
																	mediaconvert.H265RateControlModeQvbr,
																}, false),
															},
															"sample_adaptive_offset_filter_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265SampleAdaptiveOffsetFilterModeDefault,
																	mediaconvert.H265SampleAdaptiveOffsetFilterModeAdaptive,
																	mediaconvert.H265SampleAdaptiveOffsetFilterModeOff,
																}, false),
															},
															"scene_change_detect": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265SceneChangeDetectDisabled,
																	mediaconvert.H265SceneChangeDetectEnabled,
																	mediaconvert.H265SceneChangeDetectTransitionDetection,
																}, false),
															},
															"slices": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"slow_pal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265SlowPalDisabled,
																	mediaconvert.H265SlowPalEnabled,
																}, false),
																Default: mediaconvert.H265SlowPalDisabled,
															},
															"spatial_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265SpatialAdaptiveQuantizationDisabled,
																	mediaconvert.H265SpatialAdaptiveQuantizationEnabled,
																}, false),
																Default: mediaconvert.H265SpatialAdaptiveQuantizationEnabled,
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265TelecineNone,
																	mediaconvert.H265TelecineSoft,
																	mediaconvert.H265TelecineHard,
																}, false),
															},
															"temporal_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265TemporalAdaptiveQuantizationDisabled,
																	mediaconvert.H265TemporalAdaptiveQuantizationEnabled,
																}, false),
															},
															"temporal_ids": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265TemporalIdsDisabled,
																	mediaconvert.H265TemporalIdsEnabled,
																}, false),
															},
															"tiles": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265TilesDisabled,
																	mediaconvert.H265TilesEnabled,
																}, false),
															},
															"unregistered_sei_timecode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265UnregisteredSeiTimecodeDisabled,
																	mediaconvert.H265UnregisteredSeiTimecodeEnabled,
																}, false),
															},
															"write_mp4_packaging_type": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.H265WriteMp4PackagingTypeHvc1,
																	mediaconvert.H265WriteMp4PackagingTypeHev1,
																}, false),
															},
														},
													},
												},
												"mpeg2_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2AdaptiveQuantizationOff,
																	mediaconvert.Mpeg2AdaptiveQuantizationLow,
																	mediaconvert.Mpeg2AdaptiveQuantizationMedium,
																	mediaconvert.Mpeg2AdaptiveQuantizationHigh,
																}, false),
															},
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"codec_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2CodecLevelAuto,
																	mediaconvert.Mpeg2CodecLevelLow,
																	mediaconvert.Mpeg2CodecLevelMain,
																	mediaconvert.Mpeg2CodecLevelHigh1440,
																	mediaconvert.Mpeg2CodecLevelHigh,
																}, false),
															},
															"codec_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2CodecProfileMain,
																	mediaconvert.Mpeg2CodecProfileProfile422,
																}, false),
															},
															"dynamic_sub_gop": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2DynamicSubGopAdaptive,
																	mediaconvert.Mpeg2DynamicSubGopStatic,
																}, false),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2FramerateControlInitializeFromSource,
																	mediaconvert.Mpeg2FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.Mpeg2FramerateConversionAlgorithmInterpolate,
																	mediaconvert.Mpeg2FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(24),
															},
															"gop_closed_cadence": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"gop_size": {
																Type:         schema.TypeFloat,
																Optional:     true,
																ValidateFunc: validation.FloatAtLeast(0),
															},
															"gop_size_units": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2GopSizeUnitsFrames,
																	mediaconvert.Mpeg2GopSizeUnitsSeconds,
																}, false),
															},
															"hrd_buffer_initial_fill_percentage": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"hrd_buffer_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2InterlaceModeProgressive,
																	mediaconvert.Mpeg2InterlaceModeTopField,
																	mediaconvert.Mpeg2InterlaceModeBottomField,
																	mediaconvert.Mpeg2InterlaceModeFollowTopField,
																	mediaconvert.Mpeg2InterlaceModeFollowBottomField,
																}, false),
																Default: mediaconvert.Mpeg2InterlaceModeProgressive,
															},
															"intra_dc_precision": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2IntraDcPrecisionAuto,
																	mediaconvert.Mpeg2IntraDcPrecisionIntraDcPrecision8,
																	mediaconvert.Mpeg2IntraDcPrecisionIntraDcPrecision9,
																	mediaconvert.Mpeg2IntraDcPrecisionIntraDcPrecision10,
																	mediaconvert.Mpeg2IntraDcPrecisionIntraDcPrecision11,
																}, false),
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"min_i_interval": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"number_b_frames_between_reference_frames": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2ParControlInitializeFromSource,
																	mediaconvert.Mpeg2ParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality_tuning_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2QualityTuningLevelSinglePass,
																	mediaconvert.Mpeg2QualityTuningLevelMultiPass,
																}, false),
																Default: mediaconvert.Mpeg2QualityTuningLevelSinglePass,
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2RateControlModeVbr,
																	mediaconvert.Mpeg2RateControlModeCbr,
																}, false),
															},
															"scene_change_detect": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2SceneChangeDetectDisabled,
																	mediaconvert.Mpeg2SceneChangeDetectEnabled,
																}, false),
															},
															"slowpal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2SlowPalDisabled,
																	mediaconvert.Mpeg2SlowPalEnabled,
																}, false),
																Default: mediaconvert.Mpeg2SlowPalDisabled,
															},
															"softness": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntBetween(17, 128),
															},
															"spatial_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2SpatialAdaptiveQuantizationDisabled,
																	mediaconvert.Mpeg2SpatialAdaptiveQuantizationEnabled,
																}, false),
															},
															"syntax": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2SyntaxDefault,
																	mediaconvert.Mpeg2SyntaxD10,
																}, false),
																Default: mediaconvert.Mpeg2SyntaxDefault,
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2TelecineNone,
																	mediaconvert.Mpeg2TelecineSoft,
																	mediaconvert.Mpeg2TelecineHard,
																}, false),
																Default: mediaconvert.Mpeg2TelecineNone,
															},
															"temporal_adaptive_quantization": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Mpeg2TemporalAdaptiveQuantizationDisabled,
																	mediaconvert.Mpeg2TemporalAdaptiveQuantizationEnabled,
																}, false),
																Default: mediaconvert.Mpeg2TemporalAdaptiveQuantizationEnabled,
															},
														},
													},
												},
												"prores_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"codec_profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresCodecProfileAppleProres422,
																	mediaconvert.ProresCodecProfileAppleProres422Hq,
																	mediaconvert.ProresCodecProfileAppleProres422Lt,
																	mediaconvert.ProresCodecProfileAppleProres422Proxy,
																}, false),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresFramerateControlInitializeFromSource,
																	mediaconvert.ProresFramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresFramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.ProresFramerateConversionAlgorithmInterpolate,
																	mediaconvert.ProresFramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresInterlaceModeProgressive,
																	mediaconvert.ProresInterlaceModeTopField,
																	mediaconvert.ProresInterlaceModeBottomField,
																	mediaconvert.ProresInterlaceModeFollowTopField,
																	mediaconvert.ProresInterlaceModeFollowBottomField,
																}, false),
																Default: mediaconvert.ProresInterlaceModeProgressive,
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresParControlInitializeFromSource,
																	mediaconvert.ProresParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"slow_pal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresSlowPalDisabled,
																	mediaconvert.ProresSlowPalEnabled,
																}, false),
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ProresTelecineNone,
																	mediaconvert.ProresTelecineHard,
																}, false),
																Default: mediaconvert.ProresTelecineNone,
															},
														},
													},
												},
												"vc3_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3FramerateControlInitializeFromSource,
																	mediaconvert.Vc3FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.Vc3FramerateConversionAlgorithmInterpolate,
																	mediaconvert.Vc3FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(24),
															},
															"interlace_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3InterlaceModeInterlaced,
																	mediaconvert.Vc3InterlaceModeProgressive,
																}, false),
															},
															"slowpal": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3SlowPalDisabled,
																	mediaconvert.Vc3SlowPalEnabled,
																}, false),
															},
															"telecine": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3TelecineNone,
																	mediaconvert.Vc3TelecineHard,
																}, false),
															},
															"vc3_class": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vc3ClassClass1458bit,
																	mediaconvert.Vc3ClassClass2208bit,
																	mediaconvert.Vc3ClassClass22010bit,
																}, false),
															},
														},
													},
												},
												"vp8_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp8FramerateControlInitializeFromSource,
																	mediaconvert.Vp8FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp8FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.Vp8FramerateConversionAlgorithmInterpolate,
																	mediaconvert.Vp8FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"gop_size": {
																Type:         schema.TypeFloat,
																Optional:     true,
																ValidateFunc: validation.FloatAtLeast(0),
															},
															"hrd_buffer_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp8ParControlInitializeFromSource,
																	mediaconvert.Vp8ParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality_tuning_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp8QualityTuningLevelMultiPass,
																	mediaconvert.Vp8QualityTuningLevelMultiPassHq,
																}, false),
																Default: mediaconvert.Vp8QualityTuningLevelMultiPass,
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp8RateControlModeVbr,
																}, false),
															},
														},
													},
												},
												"vp9_settings": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"framerate_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp9FramerateControlInitializeFromSource,
																	mediaconvert.Vp9FramerateControlSpecified,
																}, false),
															},
															"framerate_conversion_algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp9FramerateConversionAlgorithmDuplicateDrop,
																	mediaconvert.Vp9FramerateConversionAlgorithmInterpolate,
																	mediaconvert.Vp9FramerateConversionAlgorithmFrameformer,
																}, false),
															},
															"framerate_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"framerate_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"gop_size": {
																Type:         schema.TypeFloat,
																Optional:     true,
																ValidateFunc: validation.FloatAtLeast(0),
															},
															"hrd_buffer_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"max_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1000),
															},
															"par_control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp9ParControlInitializeFromSource,
																	mediaconvert.Vp9ParControlSpecified,
																}, false),
															},
															"par_denominator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"par_numerator": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"quality_tuning_level": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp9QualityTuningLevelMultiPass,
																	mediaconvert.Vp9QualityTuningLevelMultiPassHq,
																}, false),
																Default: mediaconvert.Vp9QualityTuningLevelMultiPass,
															},
															"rate_control_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.Vp9RateControlModeVbr,
																}, false),
															},
														},
													},
												},
											},
										},
									},
									"color_metadata": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.ColorMetadataIgnore,
											mediaconvert.ColorMetadataInsert,
										}, false),
										Default: mediaconvert.ColorMetadataInsert,
									},
									"crop": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"height": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"width": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"x": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"y": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"drop_frame_timecode": {
										Optional: true,
										Type:     schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.DropFrameTimecodeDisabled,
											mediaconvert.DropFrameTimecodeEnabled,
										}, false),
									},
									"fixed_afd": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"height": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntAtLeast(32),
									},
									"position": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"height": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"width": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"x": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"y": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"respond_to_afd": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.RespondToAfdNone,
											mediaconvert.RespondToAfdRespond,
											mediaconvert.RespondToAfdPassthrough,
										}, false),
									},
									"scaling_behavior": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.ScalingBehaviorDefault,
											mediaconvert.ScalingBehaviorStretchToOutput,
										}, false),
										Default: mediaconvert.ScalingBehaviorDefault,
									},
									"sharpness": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"timecode_insertion": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.VideoTimecodeInsertionDisabled,
											mediaconvert.VideoTimecodeInsertionPicTimingSei,
										}, false),
										Default: mediaconvert.VideoTimecodeInsertionDisabled,
									},
									"video_preprocessors": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"color_corrector": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"brightness": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"color_space_conversion": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ColorSpaceConversionNone,
																	mediaconvert.ColorSpaceConversionForce601,
																	mediaconvert.ColorSpaceConversionForce709,
																	mediaconvert.ColorSpaceConversionForceHdr10,
																	mediaconvert.ColorSpaceConversionForceHlg2020,
																}, false),
															},
															"contrast": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"hdr10_metadata": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"blue_primary_x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"blue_primary_y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"green_primary_x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"green_primary_y": {
																			Type:     schema.TypeInt,
																			Optional: true},
																		"max_content_light_level": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"max_frame_average_light_level": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"max_luminance": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"min_luminance": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"red_primary_x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"red_primary_y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"white_point_x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"white_point_y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"hue": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"saturation": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
														},
													},
												},
												"deinterlacer": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"algorithm": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DeinterlaceAlgorithmInterpolate,
																	mediaconvert.DeinterlaceAlgorithmInterpolateTicker,
																	mediaconvert.DeinterlaceAlgorithmBlend,
																	mediaconvert.DeinterlaceAlgorithmBlendTicker,
																}, false),
															},
															"control": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DeinterlacerControlForceAllFrames,
																	mediaconvert.DeinterlacerControlNormal,
																}, false),
															},
															"mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DeinterlacerModeDeinterlace,
																	mediaconvert.DeinterlacerModeInverseTelecine,
																	mediaconvert.DeinterlacerModeAdaptive,
																}, false),
															},
														},
													},
												},
												"dolby_vision": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l6_metadata": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"max_cll": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"max_fall": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"l6_mode": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DolbyVisionLevel6ModePassthrough,
																	mediaconvert.DolbyVisionLevel6ModeRecalculate,
																	mediaconvert.DolbyVisionLevel6ModeSpecify,
																}, false),
															},
															"profile": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DolbyVisionProfileProfile5,
																}, false),
															},
														},
													},
												},
												"image_inserter": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"insertable_image": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"duration": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"fade_in": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"fade_out": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"height": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"image_inserter_input": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringLenBetween(14, 4000),
																		},
																		"image_x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"image_y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"layer": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"opacity": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Default:  50,
																		},
																		"start_time": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"width": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
												"noise_reducer": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"filter": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.NoiseReducerFilterBilateral,
																	mediaconvert.NoiseReducerFilterMean,
																	mediaconvert.NoiseReducerFilterGaussian,
																	mediaconvert.NoiseReducerFilterLanczos,
																	mediaconvert.NoiseReducerFilterSharpen,
																	mediaconvert.NoiseReducerFilterConserve,
																	mediaconvert.NoiseReducerFilterSpatial,
																	mediaconvert.NoiseReducerFilterTemporal,
																}, false),
															},
															"filter_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"strength": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"spatial_filter_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"post_filter_sharpen_strength": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"speed": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"strength": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"temporal_filter_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aggressive_mode": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"post_temporal_sharpening": {
																			Type:     schema.TypeString,
																			Optional: true,
																			ValidateFunc: validation.StringInSlice([]string{
																				mediaconvert.NoiseFilterPostTemporalSharpeningDisabled,
																				mediaconvert.NoiseFilterPostTemporalSharpeningEnabled,
																				mediaconvert.NoiseFilterPostTemporalSharpeningAuto,
																			}, false),
																			Default: mediaconvert.NoiseFilterPostTemporalSharpeningAuto,
																		},
																		"speed": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"strength": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
												"partner_watermaking": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"nexguard_file_marker_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"license": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringIsNotEmpty,
																		},
																		"payload": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																		"preset": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringIsNotEmpty,
																		},
																		"strength": {
																			Type:     schema.TypeString,
																			Optional: true,
																			ValidateFunc: validation.StringInSlice([]string{
																				mediaconvert.WatermarkingStrengthLightest,
																				mediaconvert.WatermarkingStrengthLighter,
																				mediaconvert.WatermarkingStrengthDefault,
																				mediaconvert.WatermarkingStrengthStronger,
																				mediaconvert.WatermarkingStrengthStrongest,
																			}, false),
																			Default: mediaconvert.WatermarkingStrengthDefault,
																		},
																	},
																},
															},
														},
													},
												},
												"timecode_burnin": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"font_size": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(10),
															},
															"position": {
																Type:     schema.TypeString,
																Optional: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.TimecodeBurninPositionTopCenter,
																	mediaconvert.TimecodeBurninPositionTopLeft,
																	mediaconvert.TimecodeBurninPositionTopRight,
																	mediaconvert.TimecodeBurninPositionMiddleLeft,
																	mediaconvert.TimecodeBurninPositionMiddleCenter,
																	mediaconvert.TimecodeBurninPositionMiddleRight,
																	mediaconvert.TimecodeBurninPositionBottomLeft,
																	mediaconvert.TimecodeBurninPositionBottomCenter,
																	mediaconvert.TimecodeBurninPositionBottomRight,
																}, false),
															},
															"prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"width": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntAtLeast(32),
									},
								},
							},
						},
					},
				},
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAwsMediaConvertPresetCreate(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	settings := &mediaconvert.PresetSettings{}
	if attr, ok := d.GetOk("settings"); ok {
		settings = expandMediaPresetSettings(attr.([]interface{}))
	}

	input := &mediaconvert.CreatePresetInput{
		Category:    aws.String(d.Get("category").(string)),
		Description: aws.String(d.Get("description").(string)),
		Name:        aws.String(d.Get("name").(string)),
		Settings:    settings,
		Tags:        keyvaluetags.New(d.Get("tags").(map[string]interface{})).IgnoreAws().MediaconvertTags(),
	}
	log.Printf("[DEBUG] Creating MediaConvert Preset: %s", input)
	return resourceAwsMediaConvertPresetRead(d, meta)
}

func expandMediaPresetSettings(list []interface{}) *mediaconvert.PresetSettings {
	presetSettings := &mediaconvert.PresetSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	settings := list[0].(map[string]interface{})
	audioDescription := &mediaconvert.AudioDescription{}
	if v, ok := settings["audio_description"]; ok {
		presetSettings.AudioDescriptions = expandMediaConvertAudioDescription(v.([]interface{}))
	}
	if v, ok := settings["caption_description"]; ok {
		presetSettings.CaptionDescriptions = expandMediaConvertCaptionDescription(v.([]interface{}))
	}

	fmt.Println(audioDescription)
	return presetSettings
}

func expandMediaConvertCaptionDescription(list []interface{}) []*mediaconvert.CaptionDescriptionPreset {
	results := []*mediaconvert.CaptionDescriptionPreset{}
	for i := 0; i < len(list); i++ {
		captionDescriptionPreset := &mediaconvert.CaptionDescriptionPreset{}
		captionDescriptionPresetMap := list[i].(map[string]interface{})
		captionDescriptionPreset.CustomLanguageCode = aws.String(captionDescriptionPresetMap["custom_language_code"].(string))
		captionDescriptionPreset.DestinationSettings = expandMediaConvertCaptionDestinationSettings(captionDescriptionPresetMap["destination_settings"].([]interface{}))
		captionDescriptionPreset.LanguageCode = aws.String(captionDescriptionPresetMap["language_code"].(string))
		captionDescriptionPreset.LanguageDescription = aws.String(captionDescriptionPresetMap["language_description"].(string))
		results = append(results, captionDescriptionPreset)
	}
	return results
}

func expandMediaConvertCaptionDestinationSettings(list []interface{}) *mediaconvert.CaptionDestinationSettings {
	result := &mediaconvert.CaptionDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	captionDestinationSettingsMap := list[0].(map[string]interface{})
	// captionDestinationSettings.BurninDestinationSettings
	// captionDestinationSettings.DestinationType
	// captionDestinationSettings.DvbSubDestinationSettings
	// captionDestinationSettings.EmbeddedDestinationSettings
	// captionDestinationSettings.ImscDestinationSettings
	// captionDestinationSettings.SccDestinationSettings
	result.TeletextDestinationSettings = expandMediaConvertTeletextDestinationSettings(captionDestinationSettingsMap["teletext_destination_settings"].([]interface{}))
	result.TtmlDestinationSettings = expandMediaConvertTtmlDestinationSettings(captionDestinationSettingsMap["ttml_destination_settings"].([]interface{}))
	return result
}

func expandMediaConvertSccDestinationSettings(list []interface{}) *mediaconvert.SccDestinationSettings {
	result := &mediaconvert.SccDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	sccDestinationSettingsMap := list[0].(map[string]interface{})
	result.Framerate = aws.String(sccDestinationSettingsMap["framerate"].(string))
	return result
}
func expandMediaConvertTeletextDestinationSettings(list []interface{}) *mediaconvert.TeletextDestinationSettings {
	result := &mediaconvert.TeletextDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	teletextDestinationSettingsMap := list[0].(map[string]interface{})
	result.PageNumber = aws.String(teletextDestinationSettingsMap["page_number"].(string))
	result.PageTypes = expandStringSet(teletextDestinationSettingsMap["page_types"].(*schema.Set))
	return result
}

func expandMediaConvertTtmlDestinationSettings(list []interface{}) *mediaconvert.TtmlDestinationSettings {
	result := &mediaconvert.TtmlDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	ttmlDestinationSettingsMap := list[0].(map[string]interface{})
	result.StylePassthrough = aws.String(ttmlDestinationSettingsMap["style_passthrough"].(string))
	return result
}

func expandMediaConvertAudioDescription(list []interface{}) []*mediaconvert.AudioDescription {
	results := []*mediaconvert.AudioDescription{}
	for i := 0; i < len(list); i++ {
		audioDescription := &mediaconvert.AudioDescription{}
		audioDescriptionMap := list[i].(map[string]interface{})
		audioDescription.AudioSourceName = aws.String(audioDescriptionMap["audio_source_name"].(string))
		audioDescription.AudioType = aws.Int64(int64(audioDescriptionMap["audio_type"].(int)))
		audioDescription.AudioTypeControl = aws.String(audioDescriptionMap["audio_type_control"].(string))
		audioDescription.CustomLanguageCode = aws.String(audioDescriptionMap["custom_language_code"].(string))
		audioDescription.LanguageCode = aws.String(audioDescriptionMap["language_code"].(string))
		audioDescription.LanguageCodeControl = aws.String(audioDescriptionMap["language_code_control"].(string))
		audioDescription.StreamName = aws.String(audioDescriptionMap["stream_name"].(string))
		audioDescription.AudioChannelTaggingSettings = expandMediaConvertAudioChannelTagging(audioDescriptionMap["audio_channel_tagging_settings"].([]interface{}))
		audioDescription.AudioNormalizationSettings = expandMediaConvertAudioNormalizationSettings(audioDescriptionMap["audio_normalization_settings"].([]interface{}))
		audioDescription.CodecSettings = expandMediaConvertCodecSettings(audioDescriptionMap["codec_settings"].([]interface{}))
		audioDescription.RemixSettings = expandMediaConvertRemixSettings(audioDescriptionMap["remix_settings"].([]interface{}))
		results = append(results, audioDescription)
	}
	return results

}

func expandMediaConvertRemixSettings(list []interface{}) *mediaconvert.RemixSettings {
	result := &mediaconvert.RemixSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	remixSettings := list[0].(map[string]interface{})
	result.ChannelMapping = expandMediaConvertChannelMapping(remixSettings["channel_mapping"].([]interface{}))
	result.ChannelsIn = aws.Int64(int64(remixSettings["channels_in"].(int)))
	result.ChannelsOut = aws.Int64(int64(remixSettings["channels_out"].(int)))
	return result
}

func expandMediaConvertChannelMapping(list []interface{}) *mediaconvert.ChannelMapping {
	channelMapping := list[0].(map[string]interface{})
	result := &mediaconvert.ChannelMapping{}
	//TODO: Test channel mapping loading

	return result
}

func expandMediaConvertCodecSettings(list []interface{}) *mediaconvert.AudioCodecSettings {
	result := &mediaconvert.AudioCodecSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	codecSettings := list[0].(map[string]interface{})
	result.Codec = aws.String(codecSettings["codec"].(string))
	result.AacSettings = expandMediaConvertAacSettings(codecSettings["aac_settings"].([]interface{}))
	result.Ac3Settings = expandMediaConvertAc3Settings(codecSettings["ac3_settings"].([]interface{}))
	result.AiffSettings = expandMediaConvertAiffSettings(codecSettings["aiff_settings"].([]interface{}))
	result.Eac3AtmosSettings = expandMediaConvertEac3AtmosSettings(codecSettings["eac3_atmos_settings"].([]interface{}))
	result.Eac3Settings = expandMediaConvertEac3Settings(codecSettings["eac3_settings"].([]interface{}))
	result.Mp2Settings = expandMediaConvertMp2Settings(codecSettings["mp2_settings"].([]interface{}))
	result.Mp3Settings = expandMediaConvertMp3Settings(codecSettings["mp3_settings"].([]interface{}))
	result.OpusSettings = expandMediaConvertOpusSettings(codecSettings["opus_settings"].([]interface{}))
	result.VorbisSettings = expandMediaConvertVorbisSettings(codecSettings["vorbis_settings"].([]interface{}))
	result.WavSettings = expandMediaConvertWavSettings(codecSettings["wav_settings"].([]interface{}))
	return result
}

func expandMediaConvertWavSettings(list []interface{}) *mediaconvert.WavSettings {
	result := &mediaconvert.WavSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	wavSettings := list[0].(map[string]interface{})
	result.BitDepth = aws.Int64(int64(wavSettings["bitdepth"].(int)))
	result.Channels = aws.Int64(int64(wavSettings["channels"].(int)))
	result.Format = aws.String(wavSettings["format"].(string))
	result.SampleRate = aws.Int64(int64(wavSettings["sample_rate"].(int)))
	return result
}
func expandMediaConvertVorbisSettings(list []interface{}) *mediaconvert.VorbisSettings {
	result := &mediaconvert.VorbisSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	vorbisSettings := list[0].(map[string]interface{})
	result.Channels = aws.Int64(int64(vorbisSettings["channels"].(int)))
	result.SampleRate = aws.Int64(int64(vorbisSettings["sample_rate"].(int)))
	result.VbrQuality = aws.Int64(int64(vorbisSettings["vbr_quality"].(int)))
	return result
}

func expandMediaConvertOpusSettings(list []interface{}) *mediaconvert.OpusSettings {
	result := &mediaconvert.OpusSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	opusSettings := list[0].(map[string]interface{})
	result.Bitrate = aws.Int64(int64(opusSettings["bitrate"].(int)))
	result.Channels = aws.Int64(int64(opusSettings["channels"].(int)))
	result.SampleRate = aws.Int64(int64(opusSettings["sample_rate"].(int)))
	return result
}

func expandMediaConvertMp3Settings(list []interface{}) *mediaconvert.Mp3Settings {
	result := &mediaconvert.Mp3Settings{}
	if list == nil || len(list) == 0 {
		return result
	}
	mp3Settings := list[0].(map[string]interface{})
	result.Bitrate = aws.Int64(int64(mp3Settings["bitrate"].(int)))
	result.Channels = aws.Int64(int64(mp3Settings["channels"].(int)))
	result.RateControlMode = aws.String(mp3Settings["rate_control_mode"].(string))
	result.SampleRate = aws.Int64(int64(mp3Settings["sample_rate"].(int)))
	result.VbrQuality = aws.Int64(int64(mp3Settings["vbr_quality"].(int)))
	return result
}

func expandMediaConvertMp2Settings(list []interface{}) *mediaconvert.Mp2Settings {
	result := &mediaconvert.Mp2Settings{}
	if list == nil || len(list) == 0 {
		return result
	}
	mp2Settings := list[0].(map[string]interface{})
	result.Bitrate = aws.Int64(int64(mp2Settings["bitrate"].(int)))
	result.Channels = aws.Int64(int64(mp2Settings["channels"].(int)))
	result.SampleRate = aws.Int64(int64(mp2Settings["sample_rate"].(int)))
	return result
}

func expandMediaConvertEac3Settings(list []interface{}) *mediaconvert.Eac3Settings {
	result := &mediaconvert.Eac3Settings{}
	if list == nil || len(list) == 0 {
		return result
	}
	eac3Settings := list[0].(map[string]interface{})
	result.AttenuationControl = aws.String(eac3Settings["attenuation_control"].(string))
	result.Bitrate = aws.Int64(int64(eac3Settings["bitrate"].(int)))
	result.BitstreamMode = aws.String(eac3Settings["bitstream_mode"].(string))
	result.CodingMode = aws.String(eac3Settings["coding_mode"].(string))
	result.DcFilter = aws.String(eac3Settings["dc_filter"].(string))
	result.Dialnorm = aws.Int64(int64(eac3Settings["dialnorm"].(int)))
	result.DynamicRangeCompressionLine = aws.String(eac3Settings["dynamic_range_compression_line"].(string))
	result.DynamicRangeCompressionRf = aws.String(eac3Settings["dynamic_range_compression_rf"].(string))
	result.LfeControl = aws.String(eac3Settings["lfe_control"].(string))
	result.LfeFilter = aws.String(eac3Settings["lfe_filter"].(string))
	result.LoRoCenterMixLevel = aws.Float64(float64(eac3Settings["lo_ro_center_mix_level"].(float32)))
	result.LoRoSurroundMixLevel = aws.Float64(float64(eac3Settings["lo_ro_surround_mix_level"].(float32)))
	result.LtRtCenterMixLevel = aws.Float64(float64(eac3Settings["lt_rt_center_mix_level"].(float32)))
	result.LtRtSurroundMixLevel = aws.Float64(float64(eac3Settings["lt_rt_surround_mix_level"].(float32)))
	result.MetadataControl = aws.String(eac3Settings["metadata_control"].(string))
	result.PassthroughControl = aws.String(eac3Settings["passthrough_control"].(string))
	result.PhaseControl = aws.String(eac3Settings["phase_control"].(string))
	result.SampleRate = aws.Int64(int64(eac3Settings["sample_rate"].(int)))
	result.StereoDownmix = aws.String(eac3Settings["stereo_downmix"].(string))
	result.SurroundExMode = aws.String(eac3Settings["surround_ex_mode"].(string))
	result.SurroundMode = aws.String(eac3Settings["surround_mode"].(string))
	return result
}

func expandMediaConvertEac3AtmosSettings(list []interface{}) *mediaconvert.Eac3AtmosSettings {
	result := &mediaconvert.Eac3AtmosSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	eac3AtmosSettings := list[0].(map[string]interface{})
	result.Bitrate = aws.Int64(int64(eac3AtmosSettings["bitrate"].(int)))
	result.BitstreamMode = aws.String(eac3AtmosSettings["bitstream_mode"].(string))
	result.CodingMode = aws.String(eac3AtmosSettings["coding_mode"].(string))
	result.DialogueIntelligence = aws.String(eac3AtmosSettings["dialogue_intelligence"].(string))
	result.DynamicRangeCompressionLine = aws.String(eac3AtmosSettings["dynamic_range_compression_line"].(string))
	result.DynamicRangeCompressionRf = aws.String(eac3AtmosSettings["dynamic_range_compression_rf"].(string))
	result.LoRoCenterMixLevel = aws.Float64(float64(eac3AtmosSettings["lo_ro_center_mix_level"].(float32)))
	result.LoRoSurroundMixLevel = aws.Float64(float64(eac3AtmosSettings["lo_ro_surround_mix_level"].(float32)))
	result.LtRtCenterMixLevel = aws.Float64(float64(eac3AtmosSettings["lt_rt_center_mix_level"].(float32)))
	result.LtRtSurroundMixLevel = aws.Float64(float64(eac3AtmosSettings["lt_rt_surround_mix_level"].(float32)))
	result.MeteringMode = aws.String(eac3AtmosSettings["metering_mode"].(string))
	result.SampleRate = aws.Int64(int64(eac3AtmosSettings["sample_rate"].(int)))
	result.SpeechThreshold = aws.Int64(int64(eac3AtmosSettings["speech_threshold"].(int)))
	result.StereoDownmix = aws.String(eac3AtmosSettings["stereo_downmix"].(string))
	result.SurroundExMode = aws.String(eac3AtmosSettings["surround_ex_mode"].(string))
	return result
}

func expandMediaConvertAiffSettings(list []interface{}) *mediaconvert.AiffSettings {
	result := &mediaconvert.AiffSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	aiffSettings := list[0].(map[string]interface{})
	result.BitDepth = aws.Int64(int64(aiffSettings["bitdepth"].(int)))
	result.Channels = aws.Int64(int64(aiffSettings["channels"].(int)))
	result.SampleRate = aws.Int64(int64(aiffSettings["sample_rate"].(int)))
	return result
}
func expandMediaConvertAc3Settings(list []interface{}) *mediaconvert.Ac3Settings {
	result := &mediaconvert.Ac3Settings{}
	if list == nil || len(list) == 0 {
		return result
	}
	ac3Settings := list[0].(map[string]interface{})
	result.Bitrate = aws.Int64(int64(ac3Settings["bitrate"].(int)))
	result.BitstreamMode = aws.String(ac3Settings["bitstream_mode"].(string))
	result.CodingMode = aws.String(ac3Settings["coding_mode"].(string))
	result.Dialnorm = aws.Int64(int64(ac3Settings["dialnorm"].(int)))
	result.DynamicRangeCompressionProfile = aws.String(ac3Settings["dynamic_range_compression_profile"].(string))
	result.LfeFilter = aws.String(ac3Settings["lfe_filter"].(string))
	result.MetadataControl = aws.String(ac3Settings["metadata_control"].(string))
	result.SampleRate = aws.Int64(int64(ac3Settings["sample_rate"].(int)))
	return result
}

func expandMediaConvertAacSettings(list []interface{}) *mediaconvert.AacSettings {
	result := &mediaconvert.AacSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	aacSettings := list[0].(map[string]interface{})
	result.AudioDescriptionBroadcasterMix = aws.String(aacSettings["audio_description_broadcaster_mix"].(string))
	result.Bitrate = aws.Int64(int64(aacSettings["bitrate"].(int)))
	result.CodecProfile = aws.String(aacSettings["codec_profile"].(string))
	result.CodingMode = aws.String(aacSettings["coding_mode"].(string))
	result.RateControlMode = aws.String(aacSettings["rate_control_mode"].(string))
	result.RawFormat = aws.String(aacSettings["raw_format"].(string))
	result.SampleRate = aws.Int64(int64(aacSettings["sample_rate"].(int)))
	result.Specification = aws.String(aacSettings["specification"].(string))
	result.VbrQuality = aws.String(aacSettings["vbr_quality"].(string))
	return result
}

func expandMediaConvertAudioNormalizationSettings(list []interface{}) *mediaconvert.AudioNormalizationSettings {
	result := &mediaconvert.AudioNormalizationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	audioNormalizationSettings := list[0].(map[string]interface{})
	result.Algorithm = aws.String(audioNormalizationSettings["algorithm"].(string))
	result.AlgorithmControl = aws.String(audioNormalizationSettings["algorithm_control"].(string))
	result.CorrectionGateLevel = aws.Int64(int64(audioNormalizationSettings["correction_gate_level"].(int)))
	result.LoudnessLogging = aws.String(audioNormalizationSettings["loudness_logging"].(string))
	result.PeakCalculation = aws.String(audioNormalizationSettings["peak_calculation"].(string))
	result.TargetLkfs = aws.Float64(float64(audioNormalizationSettings["correction_gate_level"].(float32)))
	return result
}

func expandMediaConvertAudioChannelTagging(list []interface{}) *mediaconvert.AudioChannelTaggingSettings {
	result := &mediaconvert.AudioChannelTaggingSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	channelTags := list[0].(map[string]interface{})

	result.ChannelTag = aws.String(channelTags["channel_tag"].(string))
	return result
}

func resourceAwsMediaConvertPresetRead(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}

	//ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig
	return nil
}

func resourceAwsMediaConvertPresetUpdate(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	return resourceAwsMediaConvertQueueRead(d, meta)
}

func resourceAwsMediaConvertPresetDelete(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	return nil
}
