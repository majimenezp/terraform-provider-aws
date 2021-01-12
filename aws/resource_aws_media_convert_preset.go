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
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audio_description": {
							Type:     schema.TypeList,
							MinItems: 1,
							Optional: true,
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
										MaxItems: 1,
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
										MaxItems: 1,
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
																			Type:     schema.TypeSet,
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
												"audio_buffer_model": {
													Type:     schema.TypeString,
													Optional: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsAudioBufferModelDvb,
														mediaconvert.M2tsAudioBufferModelAtsc,
													}, false),
												},
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
												"min_ebp_interval": {
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
																Type:     schema.TypeFloat,
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
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
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
	resp, err := conn.CreatePreset(input)
	if err != nil {
		return fmt.Errorf("Error creating Media Convert Preset: %s", err)
	}
	d.SetId(aws.StringValue(resp.Preset.Name))
	return resourceAwsMediaConvertPresetRead(d, meta)
}

func resourceAwsMediaConvertPresetRead(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}

	ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig
	getOpts := &mediaconvert.GetPresetInput{
		Name: aws.String(d.Id()),
	}
	resp, err := conn.GetPreset(getOpts)
	if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
		log.Printf("[WARN] Media Convert Preset (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Preset: %s", err)
	}
	d.Set("arn", resp.Preset.Arn)
	d.Set("category", resp.Preset.Category)
	d.Set("name", resp.Preset.Name)
	d.Set("description", resp.Preset.Description)
	if err := d.Set("settings", flattenMediaConvertPresetSettings(resp.Preset.Settings)); err != nil {
		return fmt.Errorf("Error setting Media Convert preset settings: %s", err)
	}

	tags, err := keyvaluetags.MediaconvertListTags(conn, aws.StringValue(resp.Preset.Arn))

	if err != nil {
		return fmt.Errorf("error listing tags for Media Convert Preset (%s): %s", d.Id(), err)
	}

	if err := d.Set("tags", tags.IgnoreAws().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return fmt.Errorf("error setting tags: %s", err)
	}
	return nil
}

func resourceAwsMediaConvertPresetUpdate(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	log.Printf("[INFO] Updating MediaConvert Preset: %s", d.Get("name").(string))
	if d.HasChanges("description", "settings") {
		updateOpts := &mediaconvert.UpdatePresetInput{
			Name: aws.String(d.Id()),
		}
		if v, ok := d.GetOk("description"); ok {
			updateOpts.Description = aws.String(v.(string))
		}

		if v, ok := d.GetOk("settings"); ok {
			presetSettings := v.([]interface{})
			updateOpts.Settings = expandMediaPresetSettings(presetSettings)
		}
		_, err = conn.UpdatePreset(updateOpts)
		if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
			log.Printf("[WARN] Media Convert Preset (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Error updating Media Convert Preset: %s", err)
		}
	}

	if d.HasChange("tags") {
		o, n := d.GetChange("tags")
		if err := keyvaluetags.MediaconvertUpdateTags(conn, d.Get("arn").(string), o, n); err != nil {
			return fmt.Errorf("error updating tags: %s", err)
		}
	}
	return resourceAwsMediaConvertPresetRead(d, meta)
}

func resourceAwsMediaConvertPresetDelete(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	delOpts := &mediaconvert.DeletePresetInput{
		Name: aws.String(d.Id()),
	}
	_, err = conn.DeletePreset(delOpts)
	if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
		return nil
	}
	if err != nil {
		return fmt.Errorf("Error deleting Media Convert Preset: %s", err)
	}
	return nil
}

func expandMediaPresetSettings(list []interface{}) *mediaconvert.PresetSettings {
	presetSettings := &mediaconvert.PresetSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	settings := list[0].(map[string]interface{})
	if v, ok := settings["audio_description"]; ok {
		presetSettings.AudioDescriptions = expandMediaConvertAudioDescription(v.([]interface{}))
	}
	if v, ok := settings["caption_description"]; ok {
		presetSettings.CaptionDescriptions = expandMediaConvertCaptionDescription(v.([]interface{}))
	}
	if v, ok := settings["container_settings"]; ok {
		presetSettings.ContainerSettings = expandMediaConvertContainerSettings(v.([]interface{}))
	}
	if v, ok := settings["video_description"]; ok {
		presetSettings.VideoDescription = expandMediaConvertVideoDescription(v.([]interface{}))
	}
	return presetSettings
}

func expandMediaConvertVideoDescription(list []interface{}) *mediaconvert.VideoDescription {
	result := &mediaconvert.VideoDescription{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["afd_signaling"].(string); ok && v != "" {
		result.AfdSignaling = aws.String(v)
	}
	if v, ok := tfMap["anti_alias"].(string); ok && v != "" {
		result.AntiAlias = aws.String(v)
	}
	if v, ok := tfMap["codec_settings"]; ok {
		result.CodecSettings = expandMediaConvertVideoCodecSettings(v.([]interface{}))
	}
	if v, ok := tfMap["color_metadata"].(string); ok && v != "" {
		result.ColorMetadata = aws.String(v)
	}
	if v, ok := tfMap["crop"]; ok {
		result.Crop = expandMediaConvertRectangle(v.([]interface{}))
	}
	if v, ok := tfMap["drop_frame_timecode"].(string); ok && v != "" {
		result.DropFrameTimecode = aws.String(v)
	}
	if v, ok := tfMap["fixed_afd"].(int); ok && v != 0 {
		result.FixedAfd = aws.Int64(int64(v))
	}
	if v, ok := tfMap["height"].(int); ok && v != 0 {
		result.Height = aws.Int64(int64(v))
	}
	if v, ok := tfMap["position"]; ok {
		result.Position = expandMediaConvertRectangle(v.([]interface{}))
	}
	if v, ok := tfMap["respond_to_afd"].(string); ok && v != "" {
		result.RespondToAfd = aws.String(v)
	}
	if v, ok := tfMap["scaling_behavior"].(string); ok && v != "" {
		result.ScalingBehavior = aws.String(v)
	}
	if v, ok := tfMap["sharpness"].(int); ok {
		result.Sharpness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["timecode_insertion"].(string); ok && v != "" {
		result.TimecodeInsertion = aws.String(v)
	}
	if v, ok := tfMap["video_preprocessors"]; ok {
		result.VideoPreprocessors = expandMediaConvertVideoPreprocessor(v.([]interface{}))
	}
	if v, ok := tfMap["width"].(int); ok && v != 0 {
		result.Width = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertVideoPreprocessor(list []interface{}) *mediaconvert.VideoPreprocessor {
	result := &mediaconvert.VideoPreprocessor{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["color_corrector"]; ok {
		result.ColorCorrector = expandMediaConvertColorCorrector(v.([]interface{}))
	}
	if v, ok := tfMap["deinterlacer"]; ok {
		result.Deinterlacer = expandMediaConvertDeinterlacer(v.([]interface{}))
	}
	if v, ok := tfMap["dolby_vision"]; ok {
		result.DolbyVision = expandMediaConvertDolbyVision(v.([]interface{}))
	}
	if v, ok := tfMap["image_inserter"]; ok {
		result.ImageInserter = expandMediaConvertImageInserter(v.([]interface{}))
	}
	if v, ok := tfMap["image_inserter"]; ok {
		result.ImageInserter = expandMediaConvertImageInserter(v.([]interface{}))
	}
	if v, ok := tfMap["noise_reducer"]; ok {
		result.NoiseReducer = expandMediaConvertNoiseReducer(v.([]interface{}))
	}
	if v, ok := tfMap["partner_watermaking"]; ok {
		result.PartnerWatermarking = expandMediaConvertPartnerWatermarking(v.([]interface{}))
	}
	if v, ok := tfMap["timecode_burnin"]; ok {
		result.TimecodeBurnin = expandMediaConvertTimecodeBurnin(v.([]interface{}))
	}
	return result
}

func expandMediaConvertTimecodeBurnin(list []interface{}) *mediaconvert.TimecodeBurnin {
	result := &mediaconvert.TimecodeBurnin{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["position"].(string); ok && v != "" {
		result.Position = aws.String(v)
	}
	if v, ok := tfMap["prefix"].(string); ok && v != "" {
		result.Prefix = aws.String(v)
	}
	return result
}
func expandMediaConvertNoiseReducer(list []interface{}) *mediaconvert.NoiseReducer {
	result := &mediaconvert.NoiseReducer{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["filter"].(string); ok && v != "" {
		result.Filter = aws.String(v)
	}
	if v, ok := tfMap["filter_settings"]; ok {
		result.FilterSettings = expandMediaConvertNoiseReducerFilterSettings(v.([]interface{}))
	}
	if v, ok := tfMap["spatial_filter_settings"]; ok {
		result.SpatialFilterSettings = expandMediaConvertNoiseReducerSpatialFilterSettings(v.([]interface{}))
	}
	if v, ok := tfMap["temporal_filter_settings"]; ok {
		result.TemporalFilterSettings = expandMediaConvertNoiseReducerTemporalFilterSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertPartnerWatermarking(list []interface{}) *mediaconvert.PartnerWatermarking {
	result := &mediaconvert.PartnerWatermarking{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["nexguard_file_marker_settings"]; ok {
		result.NexguardFileMarkerSettings = expandMediaConvertNexGuardFileMarkerSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertNexGuardFileMarkerSettings(list []interface{}) *mediaconvert.NexGuardFileMarkerSettings {
	result := &mediaconvert.NexGuardFileMarkerSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["license"].(string); ok && v != "" {
		result.License = aws.String(v)
	}
	if v, ok := tfMap["payload"].(int); ok {
		result.Payload = aws.Int64(int64(v))
	}
	if v, ok := tfMap["preset"].(string); ok && v != "" {
		result.Preset = aws.String(v)
	}
	if v, ok := tfMap["strength"].(string); ok && v != "" {
		result.Strength = aws.String(v)
	}
	return result
}
func expandMediaConvertNoiseReducerTemporalFilterSettings(list []interface{}) *mediaconvert.NoiseReducerTemporalFilterSettings {
	result := &mediaconvert.NoiseReducerTemporalFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["aggressive_mode"].(int); ok {
		result.AggressiveMode = aws.Int64(int64(v))
	}
	if v, ok := tfMap["post_temporal_sharpening"].(string); ok && v != "" {
		result.PostTemporalSharpening = aws.String(v)
	}
	if v, ok := tfMap["speed"].(int); ok {
		result.Speed = aws.Int64(int64(v))
	}
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerSpatialFilterSettings(list []interface{}) *mediaconvert.NoiseReducerSpatialFilterSettings {
	result := &mediaconvert.NoiseReducerSpatialFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["post_filter_sharpen_strength"].(int); ok {
		result.PostFilterSharpenStrength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["speed"].(int); ok {
		result.Speed = aws.Int64(int64(v))
	}
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerFilterSettings(list []interface{}) *mediaconvert.NoiseReducerFilterSettings {
	result := &mediaconvert.NoiseReducerFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertImageInserter(list []interface{}) *mediaconvert.ImageInserter {
	result := &mediaconvert.ImageInserter{}
	if list == nil || len(list) == 0 {
		return result
	}
	results := []*mediaconvert.InsertableImage{}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["insertable_image"]; ok {
		l := v.([]interface{})
		for i := 0; i < len(l); i++ {
			tmp := &mediaconvert.InsertableImage{}
			tfMap := l[i].(map[string]interface{})
			if v, ok := tfMap["duration"].(int); ok {
				tmp.Duration = aws.Int64(int64(v))
			}
			if v, ok := tfMap["fade_in"].(int); ok {
				tmp.FadeIn = aws.Int64(int64(v))
			}
			if v, ok := tfMap["fade_out"].(int); ok {
				tmp.FadeOut = aws.Int64(int64(v))
			}
			if v, ok := tfMap["height"].(int); ok && v != 0 {
				tmp.Height = aws.Int64(int64(v))
			}
			if v, ok := tfMap["image_inserter_input"].(string); ok && v != "" {
				tmp.ImageInserterInput = aws.String(v)
			}
			if v, ok := tfMap["image_x"].(int); ok {
				tmp.ImageX = aws.Int64(int64(v))
			}
			if v, ok := tfMap["image_y"].(int); ok {
				tmp.ImageY = aws.Int64(int64(v))
			}
			if v, ok := tfMap["layer"].(int); ok {
				tmp.Layer = aws.Int64(int64(v))
			}
			if v, ok := tfMap["opacity"].(int); ok {
				tmp.Opacity = aws.Int64(int64(v))
			}
			if v, ok := tfMap["start_time"].(string); ok && v != "" {
				tmp.StartTime = aws.String(v)
			}
			if v, ok := tfMap["width"].(int); ok && v != 0 {
				tmp.Width = aws.Int64(int64(v))
			}
			results = append(results, tmp)
		}
		result.InsertableImages = results
	}
	return result
}

func expandMediaConvertDolbyVision(list []interface{}) *mediaconvert.DolbyVision {
	result := &mediaconvert.DolbyVision{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["l6_metadata"]; ok {
		result.L6Metadata = expandMediaConvertDolbyVisionLevel6Metadata(v.([]interface{}))
	}
	if v, ok := tfMap["l6_mode"].(string); ok && v != "" {
		result.L6Mode = aws.String(v)
	}
	if v, ok := tfMap["profile"].(string); ok && v != "" {
		result.Profile = aws.String(v)
	}
	return result
}

func expandMediaConvertDolbyVisionLevel6Metadata(list []interface{}) *mediaconvert.DolbyVisionLevel6Metadata {
	result := &mediaconvert.DolbyVisionLevel6Metadata{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_cll"].(int); ok {
		result.MaxCll = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_fall"].(int); ok {
		result.MaxFall = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDeinterlacer(list []interface{}) *mediaconvert.Deinterlacer {
	result := &mediaconvert.Deinterlacer{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["algorithm"].(string); ok && v != "" {
		result.Algorithm = aws.String(v)
	}
	if v, ok := tfMap["control"].(string); ok && v != "" {
		result.Control = aws.String(v)
	}
	if v, ok := tfMap["mode"].(string); ok && v != "" {
		result.Mode = aws.String(v)
	}
	return result
}

func expandMediaConvertColorCorrector(list []interface{}) *mediaconvert.ColorCorrector {
	result := &mediaconvert.ColorCorrector{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["brightness"].(int); ok {
		result.Brightness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["color_space_conversion"].(string); ok && v != "" {
		result.ColorSpaceConversion = aws.String(v)
	}
	if v, ok := tfMap["contrast"].(int); ok {
		result.Contrast = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hdr10_metadata"]; ok {
		result.Hdr10Metadata = expandMediaConvertHdr10Metadata(v.([]interface{}))
	}
	if v, ok := tfMap["hue"].(int); ok {
		result.Hue = aws.Int64(int64(v))
	}
	if v, ok := tfMap["saturation"].(int); ok {
		result.Saturation = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertHdr10Metadata(list []interface{}) *mediaconvert.Hdr10Metadata {
	result := &mediaconvert.Hdr10Metadata{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["blue_primary_x"].(int); ok {
		result.BluePrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["blue_primary_y"].(int); ok {
		result.BluePrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["green_primary_x"].(int); ok {
		result.GreenPrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["green_primary_y"].(int); ok {
		result.GreenPrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_content_light_level"].(int); ok {
		result.MaxContentLightLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_frame_average_light_level"].(int); ok {
		result.MaxFrameAverageLightLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_luminance"].(int); ok {
		result.MaxLuminance = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_luminance"].(int); ok {
		result.MinLuminance = aws.Int64(int64(v))
	}
	if v, ok := tfMap["red_primary_x"].(int); ok {
		result.RedPrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["red_primary_y"].(int); ok {
		result.RedPrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["white_point_x"].(int); ok {
		result.WhitePointX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["white_point_y"].(int); ok {
		result.WhitePointY = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertRectangle(list []interface{}) *mediaconvert.Rectangle {
	result := &mediaconvert.Rectangle{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["height"].(int); ok {
		result.Height = aws.Int64(int64(v))
	}
	if v, ok := tfMap["width"].(int); ok {
		result.Width = aws.Int64(int64(v))
	}
	if v, ok := tfMap["x"].(int); ok {
		result.X = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y"].(int); ok {
		result.Y = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertVideoCodecSettings(list []interface{}) *mediaconvert.VideoCodecSettings {
	result := &mediaconvert.VideoCodecSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["av1_settings"]; ok {
		result.Av1Settings = expandMediaConvertAv1Settings(v.([]interface{}))
	}
	if v, ok := tfMap["avc_intra_settings"]; ok {
		result.AvcIntraSettings = expandMediaConvertAvcIntraSettings(v.([]interface{}))
	}
	if v, ok := tfMap["codec"].(string); ok && v != "" {
		result.Codec = aws.String(v)
	}
	if v, ok := tfMap["frame_capture_settings"]; ok {
		result.FrameCaptureSettings = expandMediaConvertFrameCaptureSettings(v.([]interface{}))
	}
	if v, ok := tfMap["h264_settings"]; ok {
		result.H264Settings = expandMediaConvertH264Settings(v.([]interface{}))
	}
	if v, ok := tfMap["h265_settings"]; ok {
		result.H265Settings = expandMediaConvertH265Settings(v.([]interface{}))
	}
	if v, ok := tfMap["mpeg2_settings"]; ok {
		result.Mpeg2Settings = expandMediaConvertMpeg2Settings(v.([]interface{}))
	}
	if v, ok := tfMap["prores_settings "]; ok {
		result.ProresSettings = expandMediaConvertProresSettings(v.([]interface{}))
	}
	if v, ok := tfMap["vc3_settings"]; ok {
		result.Vc3Settings = expandMediaConvertVc3Settings(v.([]interface{}))
	}
	if v, ok := tfMap["vp8_settings"]; ok {
		result.Vp8Settings = expandMediaConvertVp8Settings(v.([]interface{}))
	}
	if v, ok := tfMap["vp9_settings"]; ok {
		result.Vp9Settings = expandMediaConvertVp9Settings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertProresSettings(list []interface{}) *mediaconvert.ProresSettings {
	result := &mediaconvert.ProresSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	return result
}

func expandMediaConvertVc3Settings(list []interface{}) *mediaconvert.Vc3Settings {
	result := &mediaconvert.Vc3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["vc3_class"].(string); ok && v != "" {
		result.Vc3Class = aws.String(v)
	}
	return result
}

func expandMediaConvertVp8Settings(list []interface{}) *mediaconvert.Vp8Settings {
	result := &mediaconvert.Vp8Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	return result
}

func expandMediaConvertVp9Settings(list []interface{}) *mediaconvert.Vp9Settings {
	result := &mediaconvert.Vp9Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	return result
}

func expandMediaConvertMpeg2Settings(list []interface{}) *mediaconvert.Mpeg2Settings {
	result := &mediaconvert.Mpeg2Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["intra_dc_precision"].(string); ok && v != "" {
		result.IntraDcPrecision = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["softness"].(int); ok {
		result.Softness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["syntax"].(string); ok && v != "" {
		result.Syntax = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	return result
}

func expandMediaConvertH265Settings(list []interface{}) *mediaconvert.H265Settings {
	result := &mediaconvert.H265Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["alternate_transfer_function_sei"].(string); ok && v != "" {
		result.AlternateTransferFunctionSei = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_b_reference"].(string); ok && v != "" {
		result.GopBReference = aws.String(v)
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_reference_frames"].(int); ok {
		result.NumberReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertH265QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["sample_adaptive_offset_filter_mode"].(string); ok && v != "" {
		result.SampleAdaptiveOffsetFilterMode = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["temporal_ids"].(string); ok && v != "" {
		result.TemporalIds = aws.String(v)
	}
	if v, ok := tfMap["tiles"].(string); ok && v != "" {
		result.Tiles = aws.String(v)
	}
	if v, ok := tfMap["unregistered_sei_timecode"].(string); ok && v != "" {
		result.UnregisteredSeiTimecode = aws.String(v)
	}
	if v, ok := tfMap["write_mp4_packaging_type"].(string); ok && v != "" {
		result.WriteMp4PackagingType = aws.String(v)
	}
	return result
}

func expandMediaConvertH264Settings(list []interface{}) *mediaconvert.H264Settings {
	result := &mediaconvert.H264Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok && v != 0 {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["entropy_encoding"].(string); ok && v != "" {
		result.EntropyEncoding = aws.String(v)
	}
	if v, ok := tfMap["field_encoding"].(string); ok && v != "" {
		result.FieldEncoding = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_b_reference"].(string); ok && v != "" {
		result.GopBReference = aws.String(v)
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_reference_frames"].(int); ok {
		result.NumberReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertH264QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["repeat_pps"].(string); ok && v != "" {
		result.RepeatPps = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["softness"].(int); ok {
		result.Softness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["syntax"].(string); ok && v != "" {
		result.Syntax = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["unregistered_sei_timecode"].(string); ok && v != "" {
		result.UnregisteredSeiTimecode = aws.String(v)
	}
	return result
}

func expandMediaConvertH265QvbrSettings(list []interface{}) *mediaconvert.H265QvbrSettings {
	result := &mediaconvert.H265QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_average_bitrate"].(int); ok && v != 0 {
		result.MaxAverageBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertH264QvbrSettings(list []interface{}) *mediaconvert.H264QvbrSettings {
	result := &mediaconvert.H264QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_average_bitrate"].(int); ok && v != 0 {
		result.MaxAverageBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertFrameCaptureSettings(list []interface{}) *mediaconvert.FrameCaptureSettings {
	result := &mediaconvert.FrameCaptureSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_captures"].(int); ok {
		result.MaxCaptures = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality"].(int); ok {
		result.Quality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAvcIntraSettings(list []interface{}) *mediaconvert.AvcIntraSettings {
	result := &mediaconvert.AvcIntraSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["avc_intra_class"].(string); ok && v != "" {
		result.AvcIntraClass = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	return result
}

func expandMediaConvertAv1Settings(list []interface{}) *mediaconvert.Av1Settings {
	result := &mediaconvert.Av1Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertAv1QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	return result
}

func expandMediaConvertAv1QvbrSettings(list []interface{}) *mediaconvert.Av1QvbrSettings {
	result := &mediaconvert.Av1QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertContainerSettings(list []interface{}) *mediaconvert.ContainerSettings {
	result := &mediaconvert.ContainerSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	containerSettingsMap := list[0].(map[string]interface{})
	if v, ok := containerSettingsMap["cmfc_settings"]; ok {
		result.CmfcSettings = expandMediaConvertCmfcSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["container"].(string); ok && v != "" {
		result.Container = aws.String(v)
	}
	if v, ok := containerSettingsMap["f4v_settings"]; ok {
		result.F4vSettings = expandMediaConvertF4vSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["m2ts_settings"]; ok {
		result.M2tsSettings = expandMediaConvertM2tsSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["m3u8_settings"]; ok {
		result.M3u8Settings = expandMediaConvertM3u8Settings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mov_settings"]; ok {
		result.MovSettings = expandMediaConvertMovSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mp4_settings"]; ok {
		result.Mp4Settings = expandMediaConvertMp4Settings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mpd_settings"]; ok {
		result.MpdSettings = expandMediaConvertMpdSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mxf_settings"]; ok {
		result.MxfSettings = expandMediaConvertMxfSettings(v.([]interface{}))
	}

	return result
}

func expandMediaConvertMxfSettings(list []interface{}) *mediaconvert.MxfSettings {
	result := &mediaconvert.MxfSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["afd_signaling"].(string); ok && v != "" {
		result.AfdSignaling = aws.String(v)
	}
	if v, ok := tfMap["profile"].(string); ok && v != "" {
		result.Profile = aws.String(v)
	}
	return result
}
func expandMediaConvertMpdSettings(list []interface{}) *mediaconvert.MpdSettings {
	result := &mediaconvert.MpdSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["accessibility_caption_hints"].(string); ok && v != "" {
		result.AccessibilityCaptionHints = aws.String(v)
	}
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["caption_container_type"].(string); ok && v != "" {
		result.CaptionContainerType = aws.String(v)
	}
	if v, ok := tfMap["scte_35_esam"].(string); ok && v != "" {
		result.Scte35Esam = aws.String(v)
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	return result
}

func expandMediaConvertMp4Settings(list []interface{}) *mediaconvert.Mp4Settings {
	result := &mediaconvert.Mp4Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["cslg_atom"].(string); ok && v != "" {
		result.CslgAtom = aws.String(v)
	}
	if v, ok := tfMap["ctts_version"].(int); ok {
		result.CttsVersion = aws.Int64(int64(v))
	}
	if v, ok := tfMap["free_space_box"].(string); ok && v != "" {
		result.FreeSpaceBox = aws.String(v)
	}
	if v, ok := tfMap["moov_placement"].(string); ok && v != "" {
		result.MoovPlacement = aws.String(v)
	}
	if v, ok := tfMap["mp4_major_brand"].(string); ok && v != "" {
		result.Mp4MajorBrand = aws.String(v)
	}
	return result
}

func expandMediaConvertMovSettings(list []interface{}) *mediaconvert.MovSettings {
	result := &mediaconvert.MovSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["clap_atom"].(string); ok && v != "" {
		result.ClapAtom = aws.String(v)
	}
	if v, ok := tfMap["cslg_atom"].(string); ok && v != "" {
		result.CslgAtom = aws.String(v)
	}
	if v, ok := tfMap["mpeg2_fourcc_control"].(string); ok && v != "" {
		result.Mpeg2FourCCControl = aws.String(v)
	}
	if v, ok := tfMap["padding_control"].(string); ok && v != "" {
		result.PaddingControl = aws.String(v)
	}
	if v, ok := tfMap["reference"].(string); ok && v != "" {
		result.Reference = aws.String(v)
	}
	return result
}

func expandMediaConvertM3u8Settings(list []interface{}) *mediaconvert.M3u8Settings {
	result := &mediaconvert.M3u8Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["audio_frames_per_pes"].(int); ok {
		result.AudioFramesPerPes = aws.Int64(int64(v))
	}
	if v, ok := tfMap["audio_pids"].(*schema.Set); ok && v.Len() > 0 {
		result.AudioPids = expandInt64Set(v)
	}
	if v, ok := tfMap["nielsen_id3"].(string); ok && v != "" {
		result.NielsenId3 = aws.String(v)
	}
	if v, ok := tfMap["pat_interval"].(int); ok {
		result.PatInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pcr_control"].(string); ok && v != "" {
		result.PcrControl = aws.String(v)
	}
	if v, ok := tfMap["pcr_pid"].(int); ok {
		result.PcrPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_interval"].(int); ok {
		result.PmtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_pid"].(int); ok {
		result.PmtPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["private_metadata_pid"].(int); ok {
		result.PrivateMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["program_number"].(int); ok {
		result.ProgramNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_pid"].(int); ok {
		result.Scte35Pid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata"].(string); ok && v != "" {
		result.TimedMetadata = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata_pid"].(int); ok {
		result.TimedMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["transport_stream_id"].(int); ok {
		result.TransportStreamId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["video_pid"].(int); ok {
		result.VideoPid = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertM2tsSettings(list []interface{}) *mediaconvert.M2tsSettings {
	result := &mediaconvert.M2tsSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_buffer_model"].(string); ok && v != "" {
		result.AudioBufferModel = aws.String(v)
	}
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["audio_frames_per_pes"].(int); ok {
		result.AudioFramesPerPes = aws.Int64(int64(v))
	}
	if v, ok := tfMap["protocols"].(*schema.Set); ok && v.Len() > 0 {
		result.AudioPids = expandInt64Set(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["buffer_model"].(string); ok && v != "" {
		result.BufferModel = aws.String(v)
	}
	if v, ok := tfMap["dvb_nit_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbNitSettings = expandMediaConvertDvbNitSettings(v.List())
	}
	if v, ok := tfMap["dvb_sdt_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbSdtSettings = expandMediaConvertDvbSdtSettings(v.List())
	}
	if v, ok := tfMap["dvb_sub_pids"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbSubPids = expandInt64Set(v)
	}
	if v, ok := tfMap["dvb_tdt_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbTdtSettings = expandMediaConvertDvbTdtSettings(v.List())
	}
	if v, ok := tfMap["dvb_teletext_pid"].(int); ok {
		result.DvbTeletextPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["ebp_audio_interval"].(string); ok && v != "" {
		result.EbpAudioInterval = aws.String(tfMap["ebp_audio_interval"].(string))
	}
	if v, ok := tfMap["ebp_placement"].(string); ok && v != "" {
		result.EbpPlacement = aws.String(v)
	}
	if v, ok := tfMap["es_rate_in_pes"].(string); ok && v != "" {
		result.EsRateInPes = aws.String(v)
	}
	if v, ok := tfMap["force_ts_video_ebp_order"].(string); ok && v != "" {
		result.ForceTsVideoEbpOrder = aws.String(v)
	}
	if v, ok := tfMap["fragment_time"].(float64); ok {
		result.FragmentTime = aws.Float64(float64(v))
	}
	if v, ok := tfMap["max_pcr_interval"].(int); ok {
		result.MaxPcrInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_ebp_interval"].(int); ok {
		result.MinEbpInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["nielsen_id3"].(string); ok && v != "" {
		result.NielsenId3 = aws.String(v)
	}
	if v, ok := tfMap["null_packet_bitrate"].(float64); ok {
		result.NullPacketBitrate = aws.Float64(float64(v))
	}
	if v, ok := tfMap["pat_interval"].(int); ok {
		result.PatInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pcr_control"].(string); ok && v != "" {
		result.PcrControl = aws.String(v)
	}
	if v, ok := tfMap["pcr_pid"].(int); ok {
		result.PcrPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_interval"].(int); ok {
		result.PmtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_pid"].(int); ok {
		result.PmtPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["private_metadata_pid"].(int); ok {
		result.PrivateMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["program_number"].(int); ok {
		result.ProgramNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["rate_mode"].(string); ok && v != "" {
		result.RateMode = aws.String(v)
	}
	if v, ok := tfMap["scte_35_esam"].(*schema.Set); ok && v.Len() > 0 {
		result.Scte35Esam = expandMediaConvertM2tsScte35Esam(v.List())
	}
	if v, ok := tfMap["scte_35_pid"].(int); ok {
		result.Scte35Pid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	if v, ok := tfMap["segmentation_markers"].(string); ok && v != "" {
		result.SegmentationMarkers = aws.String(v)
	}
	if v, ok := tfMap["segmentation_style"].(string); ok && v != "" {
		result.SegmentationStyle = aws.String(v)
	}
	if v, ok := tfMap["segmentation_time"].(float64); ok {
		result.SegmentationTime = aws.Float64(float64(v))
	}
	if v, ok := tfMap["timed_metadata_pid"].(int); ok {
		result.TimedMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["transport_stream_id"].(int); ok {
		result.TransportStreamId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["video_pid"].(int); ok {
		result.VideoPid = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDvbTdtSettings(list []interface{}) *mediaconvert.DvbTdtSettings {
	result := &mediaconvert.DvbTdtSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["tdt_interval"].(int); ok {
		result.TdtInterval = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertM2tsScte35Esam(list []interface{}) *mediaconvert.M2tsScte35Esam {
	result := &mediaconvert.M2tsScte35Esam{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["scte_35_esam_pid"].(int); ok {
		result.Scte35EsamPid = aws.Int64(int64(v))
	}
	return result
}
func expandMediaConvertDvbNitSettings(list []interface{}) *mediaconvert.DvbNitSettings {
	result := &mediaconvert.DvbNitSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["network_id"].(int); ok {
		result.NetworkId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["network_name"].(string); ok && v != "" {
		result.NetworkName = aws.String(v)
	}
	if v, ok := tfMap["nit_interval"].(int); ok {
		result.NitInterval = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDvbSdtSettings(list []interface{}) *mediaconvert.DvbSdtSettings {
	result := &mediaconvert.DvbSdtSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["output_sdt"].(string); ok && v != "" {
		result.OutputSdt = aws.String(v)
	}
	if v, ok := tfMap["sdt_interval"].(int); ok {
		result.SdtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["service_name"].(string); ok && v != "" {
		result.ServiceName = aws.String(v)
	}
	if v, ok := tfMap["service_provider_name"].(string); ok && v != "" {
		result.ServiceProviderName = aws.String(v)
	}

	return result
}

func expandMediaConvertF4vSettings(list []interface{}) *mediaconvert.F4vSettings {
	result := &mediaconvert.F4vSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["moov_placement"].(string); ok && v != "" {
		result.MoovPlacement = aws.String(tfMap["moov_placement"].(string))
	}
	return result
}

func expandMediaConvertCmfcSettings(list []interface{}) *mediaconvert.CmfcSettings {
	result := &mediaconvert.CmfcSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["scte35_esam"].(string); ok && v != "" {
		result.Scte35Esam = aws.String(v)
	}
	if v, ok := tfMap["scte35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	return result
}

func expandMediaConvertCaptionDescription(list []interface{}) []*mediaconvert.CaptionDescriptionPreset {
	if list == nil || len(list) == 0 {
		return nil
	}
	results := []*mediaconvert.CaptionDescriptionPreset{}
	for i := 0; i < len(list); i++ {
		captionDescriptionPreset := &mediaconvert.CaptionDescriptionPreset{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
			captionDescriptionPreset.CustomLanguageCode = aws.String(v)
		}
		captionDescriptionPreset.DestinationSettings = expandMediaConvertCaptionDestinationSettings(tfMap["destination_settings"].([]interface{}))
		if v, ok := tfMap["language_code"].(string); ok && v != "" {
			captionDescriptionPreset.LanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_description"].(string); ok && v != "" {
			captionDescriptionPreset.LanguageDescription = aws.String(v)
		}
		results = append(results, captionDescriptionPreset)
	}
	return results
}

func expandMediaConvertCaptionDestinationSettings(list []interface{}) *mediaconvert.CaptionDestinationSettings {
	result := &mediaconvert.CaptionDestinationSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["burnin_destination_settings"]; ok {
		result.BurninDestinationSettings = expandMediaConvertBurninDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["destination_type"].(string); ok && v != "" {
		result.DestinationType = aws.String(v)
	}
	if v, ok := tfMap["dvb_sub_destination_settings"]; ok {
		result.DvbSubDestinationSettings = expandMediaConvertDvbSubDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["embedded_destination_settings "]; ok {
		result.EmbeddedDestinationSettings = expandMediaConvertEmbeddedDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["imsc_destination_settings "]; ok {
		result.ImscDestinationSettings = expandMediaConvertImscDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["scc_destination_settings "]; ok {
		result.SccDestinationSettings = expandMediaConvertSccDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["teletext_destination_settings "]; ok {
		result.TeletextDestinationSettings = expandMediaConvertTeletextDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["ttml_destination_settings "]; ok {
		result.TtmlDestinationSettings = expandMediaConvertTtmlDestinationSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertBurninDestinationSettings(list []interface{}) *mediaconvert.BurninDestinationSettings {
	result := &mediaconvert.BurninDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["alignment"].(string); ok && v != "" {
		result.Alignment = aws.String(v)
	}
	if v, ok := tfMap["background_color"].(string); ok && v != "" {
		result.BackgroundColor = aws.String(v)
	}
	if v, ok := tfMap["background_opacity"].(int); ok {
		result.BackgroundOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_color"].(string); ok && v != "" {
		result.FontColor = aws.String(v)
	}
	if v, ok := tfMap["font_opacity"].(int); ok {
		result.FontOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_resolution"].(int); ok {
		result.FontResolution = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_script"].(string); ok && v != "" {
		result.FontScript = aws.String(v)
	}
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["outline_color"].(string); ok && v != "" {
		result.OutlineColor = aws.String(v)
	}
	if v, ok := tfMap["outline_size"].(int); ok {
		result.OutlineSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_color"].(string); ok && v != "" {
		result.ShadowColor = aws.String(v)
	}
	if v, ok := tfMap["shadow_opacity"].(int); ok {
		result.ShadowOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_x_offset"].(int); ok {
		result.ShadowXOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_y_offset"].(int); ok {
		result.ShadowYOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["teletext_spacing"].(string); ok && v != "" {
		result.TeletextSpacing = aws.String(v)
	}
	if v, ok := tfMap["x_position"].(int); ok {
		result.XPosition = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y_position"].(int); ok {
		result.YPosition = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDvbSubDestinationSettings(list []interface{}) *mediaconvert.DvbSubDestinationSettings {
	result := &mediaconvert.DvbSubDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["alignment"].(string); ok && v != "" {
		result.Alignment = aws.String(v)
	}
	if v, ok := tfMap["background_color"].(string); ok && v != "" {
		result.BackgroundColor = aws.String(v)
	}
	if v, ok := tfMap["background_opacity"].(int); ok {
		result.BackgroundOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_color"].(string); ok && v != "" {
		result.FontColor = aws.String(v)
	}
	if v, ok := tfMap["font_opacity"].(int); ok {
		result.FontOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_resolution"].(int); ok {
		result.FontResolution = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_script"].(string); ok && v != "" {
		result.FontScript = aws.String(v)
	}
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["outline_color"].(string); ok && v != "" {
		result.OutlineColor = aws.String(v)
	}
	if v, ok := tfMap["outline_size"].(int); ok {
		result.OutlineSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_color"].(string); ok && v != "" {
		result.ShadowColor = aws.String(v)
	}
	if v, ok := tfMap["shadow_opacity"].(int); ok {
		result.ShadowOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_x_offset"].(int); ok {
		result.ShadowXOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_y_offset"].(int); ok {
		result.ShadowYOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["subtitling_type"].(string); ok && v != "" {
		result.SubtitlingType = aws.String(v)
	}
	if v, ok := tfMap["teletext_spacing"].(string); ok && v != "" {
		result.TeletextSpacing = aws.String(v)
	}
	if v, ok := tfMap["x_position"].(int); ok {
		result.XPosition = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y_position"].(int); ok {
		result.YPosition = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertEmbeddedDestinationSettings(list []interface{}) *mediaconvert.EmbeddedDestinationSettings {
	result := &mediaconvert.EmbeddedDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["destination_608_channel_number"].(int); ok {
		result.Destination608ChannelNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["destination_708_service_number"].(int); ok {
		result.Destination708ServiceNumber = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertImscDestinationSettings(list []interface{}) *mediaconvert.ImscDestinationSettings {
	result := &mediaconvert.ImscDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["style_passthrough"].(string); ok && v != "" {
		result.StylePassthrough = aws.String(v)
	}
	return result
}

func expandMediaConvertSccDestinationSettings(list []interface{}) *mediaconvert.SccDestinationSettings {
	result := &mediaconvert.SccDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate"].(string); ok && v != "" {
		result.Framerate = aws.String(v)
	}
	return result
}
func expandMediaConvertTeletextDestinationSettings(list []interface{}) *mediaconvert.TeletextDestinationSettings {
	result := &mediaconvert.TeletextDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["page_number"].(string); ok && v != "" {
		result.PageNumber = aws.String(v)
	}
	result.PageTypes = expandStringSet(tfMap["page_types"].(*schema.Set))
	return result
}

func expandMediaConvertTtmlDestinationSettings(list []interface{}) *mediaconvert.TtmlDestinationSettings {
	result := &mediaconvert.TtmlDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["style_passthrough"].(string); ok && v != "" {
		result.StylePassthrough = aws.String(v)
	}
	return result
}

func expandMediaConvertAudioDescription(list []interface{}) []*mediaconvert.AudioDescription {
	results := []*mediaconvert.AudioDescription{}
	if list == nil || len(list) == 0 {
		return nil
	}
	for i := 0; i < len(list); i++ {
		audioDescription := &mediaconvert.AudioDescription{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["audio_source_name"].(string); ok && v != "" {
			audioDescription.AudioSourceName = aws.String(v)
		}
		if v, ok := tfMap["audio_type"].(int); ok && v != 0 {
			audioDescription.AudioType = aws.Int64(int64(v))
		}
		if v, ok := tfMap["audio_type_control"].(string); ok && v != "" {
			audioDescription.AudioTypeControl = aws.String(v)
		}
		if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
			audioDescription.CustomLanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_code"].(string); ok && v != "" {
			audioDescription.LanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_code_control"].(string); ok && v != "" {
			audioDescription.LanguageCodeControl = aws.String(v)
		}
		if v, ok := tfMap["stream_name"].(string); ok && v != "" {
			audioDescription.StreamName = aws.String(v)
		}
		audioDescription.AudioChannelTaggingSettings = expandMediaConvertAudioChannelTagging(tfMap["audio_channel_tagging_settings"].([]interface{}))
		audioDescription.AudioNormalizationSettings = expandMediaConvertAudioNormalizationSettings(tfMap["audio_normalization_settings"].([]interface{}))
		audioDescription.CodecSettings = expandMediaConvertCodecSettings(tfMap["codec_settings"].([]interface{}))
		audioDescription.RemixSettings = expandMediaConvertRemixSettings(tfMap["remix_settings"].([]interface{}))
		results = append(results, audioDescription)
	}
	return results

}

func expandMediaConvertRemixSettings(list []interface{}) *mediaconvert.RemixSettings {
	result := &mediaconvert.RemixSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	result.ChannelMapping = expandMediaConvertChannelMapping(tfMap["channel_mapping"].([]interface{}))
	if v, ok := tfMap["channels_in"].(int); ok {
		result.ChannelsIn = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels_out"].(int); ok {
		result.ChannelsOut = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertChannelMapping(list []interface{}) *mediaconvert.ChannelMapping {
	channelMapping := list[0].(map[string]interface{})
	result := &mediaconvert.ChannelMapping{}
	outputChannels := []*mediaconvert.OutputChannelMapping{}
	l := channelMapping["output_channels"].([]interface{})
	for _, tfMapRaw := range l {
		tfMap, ok := tfMapRaw.(map[string]interface{})
		if !ok {
			continue
		}
		outputChannel := &mediaconvert.OutputChannelMapping{}
		if v, ok := tfMap["input_channels"].(*schema.Set); ok && v.Len() > 0 {
			outputChannel.InputChannels = expandInt64Set(v)
		}
		outputChannels = append(outputChannels, outputChannel)
	}
	result.OutputChannels = outputChannels
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
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitdepth"].(int); ok {
		result.BitDepth = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["format"].(string); ok && v != "" {
		result.Format = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}
func expandMediaConvertVorbisSettings(list []interface{}) *mediaconvert.VorbisSettings {
	result := &mediaconvert.VorbisSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["vbr_quality"].(int); ok {
		result.VbrQuality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertOpusSettings(list []interface{}) *mediaconvert.OpusSettings {
	result := &mediaconvert.OpusSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertMp3Settings(list []interface{}) *mediaconvert.Mp3Settings {
	result := &mediaconvert.Mp3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["vbr_quality"].(int); ok {
		result.VbrQuality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertMp2Settings(list []interface{}) *mediaconvert.Mp2Settings {
	result := &mediaconvert.Mp2Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertEac3Settings(list []interface{}) *mediaconvert.Eac3Settings {
	result := &mediaconvert.Eac3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["attenuation_control"].(string); ok && v != "" {
		result.AttenuationControl = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dc_filter"].(string); ok && v != "" {
		result.DcFilter = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Dialnorm = aws.Int64(int64(v))
	}
	if v, ok := tfMap["dynamic_range_compression_line"].(string); ok && v != "" {
		result.DynamicRangeCompressionLine = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_rf"].(string); ok && v != "" {
		result.DynamicRangeCompressionRf = aws.String(v)
	}
	if v, ok := tfMap["lfe_control"].(string); ok && v != "" {
		result.LfeControl = aws.String(v)
	}
	if v, ok := tfMap["lfe_filter"].(string); ok && v != "" {
		result.LfeFilter = aws.String(v)
	}
	if v, ok := tfMap["lo_ro_center_mix_level"].(float64); ok {
		result.LoRoCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lo_ro_surround_mix_level"].(float64); ok {
		result.LoRoSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_center_mix_level"].(float64); ok {
		result.LtRtCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_surround_mix_level"].(float64); ok {
		result.LtRtSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["metadata_control"].(string); ok && v != "" {
		result.MetadataControl = aws.String(v)
	}
	if v, ok := tfMap["passthrough_control"].(string); ok && v != "" {
		result.PassthroughControl = aws.String(v)
	}
	if v, ok := tfMap["phase_control"].(string); ok && v != "" {
		result.PhaseControl = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stereo_downmix"].(string); ok && v != "" {
		result.StereoDownmix = aws.String(v)
	}
	if v, ok := tfMap["surround_ex_mode"].(string); ok && v != "" {
		result.SurroundExMode = aws.String(v)
	}
	if v, ok := tfMap["surround_mode"].(string); ok && v != "" {
		result.SurroundMode = aws.String(v)
	}
	return result
}

func expandMediaConvertEac3AtmosSettings(list []interface{}) *mediaconvert.Eac3AtmosSettings {
	result := &mediaconvert.Eac3AtmosSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dialogue_intelligence"].(string); ok && v != "" {
		result.DialogueIntelligence = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_line"].(string); ok && v != "" {
		result.DynamicRangeCompressionLine = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_rf"].(string); ok && v != "" {
		result.DynamicRangeCompressionRf = aws.String(v)
	}
	if v, ok := tfMap["lo_ro_center_mix_level"].(float64); ok {
		result.LoRoCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lo_ro_surround_mix_level"].(float64); ok {
		result.LoRoSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_center_mix_level"].(float64); ok {
		result.LtRtCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_surround_mix_level"].(float64); ok {
		result.LtRtSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["metering_mode"].(string); ok && v != "" {
		result.MeteringMode = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["speech_threshold"].(int); ok {
		result.SpeechThreshold = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stereo_downmix"].(string); ok && v != "" {
		result.StereoDownmix = aws.String(v)
	}
	if v, ok := tfMap["surround_ex_mode"].(string); ok && v != "" {
		result.SurroundExMode = aws.String(v)
	}
	return result
}

func expandMediaConvertAiffSettings(list []interface{}) *mediaconvert.AiffSettings {
	result := &mediaconvert.AiffSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitdepth"].(int); ok {
		result.BitDepth = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}
func expandMediaConvertAc3Settings(list []interface{}) *mediaconvert.Ac3Settings {
	result := &mediaconvert.Ac3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dialnorm"].(int); ok {
		result.Dialnorm = aws.Int64(int64(v))
	}
	if v, ok := tfMap["dynamic_range_compression_profile"].(string); ok && v != "" {
		result.DynamicRangeCompressionProfile = aws.String(v)
	}
	if v, ok := tfMap["lfe_filter"].(string); ok && v != "" {
		result.LfeFilter = aws.String(v)
	}
	if v, ok := tfMap["metadata_control"].(string); ok && v != "" {
		result.MetadataControl = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAacSettings(list []interface{}) *mediaconvert.AacSettings {
	result := &mediaconvert.AacSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_description_broadcaster_mix"].(string); ok && v != "" {
		result.AudioDescriptionBroadcasterMix = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok && v != 0 {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["raw_format"].(string); ok && v != "" {
		result.RawFormat = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok && v != 0 {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["specification"].(string); ok && v != "" {
		result.Specification = aws.String(v)
	}
	if v, ok := tfMap["vbr_quality"].(string); ok && v != "" {
		result.VbrQuality = aws.String(v)
	}
	return result
}

func expandMediaConvertAudioNormalizationSettings(list []interface{}) *mediaconvert.AudioNormalizationSettings {
	result := &mediaconvert.AudioNormalizationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["algorithm"].(string); ok && v != "" {
		result.Algorithm = aws.String(v)
	}
	if v, ok := tfMap["algorithm_control"].(string); ok && v != "" {
		result.AlgorithmControl = aws.String(v)
	}
	if v, ok := tfMap["correction_gate_level"].(int); ok {
		result.CorrectionGateLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["loudness_logging"].(string); ok && v != "" {
		result.LoudnessLogging = aws.String(v)
	}
	if v, ok := tfMap["peak_calculation"].(string); ok && v != "" {
		result.PeakCalculation = aws.String(v)
	}
	if v, ok := tfMap["target_lkfs"].(float64); ok {
		result.TargetLkfs = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertAudioChannelTagging(list []interface{}) *mediaconvert.AudioChannelTaggingSettings {
	result := &mediaconvert.AudioChannelTaggingSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["channel_tag"].(string); ok && v != "" {
		result.ChannelTag = aws.String(v)
	}
	return result
}

func flattenMediaConvertPresetSettings(presetSettings *mediaconvert.PresetSettings) []interface{} {
	if presetSettings == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_description":   flattenMediaConvertAudioDescription(presetSettings.AudioDescriptions),
		"caption_description": flattenMediaConvertCaptionDescriptionPreset(presetSettings.CaptionDescriptions),
		"container_settings":  flattenMediaConvertContainerSettings(presetSettings.ContainerSettings),
		"video_description":   flattenMediaConvertVideoDescription(presetSettings.VideoDescription),
	}
	return []interface{}{m}
}

func flattenMediaConvertAudioDescription(audioDescriptions []*mediaconvert.AudioDescription) []interface{} {
	if audioDescriptions == nil {
		return []interface{}{}
	}
	audioDesc := make([]interface{}, 0, len(audioDescriptions))
	for i := 0; i < len(audioDescriptions); i++ {
		m := map[string]interface{}{
			"audio_source_name":              aws.StringValue(audioDescriptions[i].AudioSourceName),
			"audio_type":                     aws.Int64Value(audioDescriptions[i].AudioType),
			"audio_type_control":             aws.StringValue(audioDescriptions[i].AudioTypeControl),
			"custom_language_code":           aws.StringValue(audioDescriptions[i].CustomLanguageCode),
			"language_code":                  aws.StringValue(audioDescriptions[i].LanguageCode),
			"language_code_control":          aws.StringValue(audioDescriptions[i].LanguageCodeControl),
			"stream_name":                    aws.StringValue(audioDescriptions[i].StreamName),
			"audio_channel_tagging_settings": flattenMediaConvertAudioChannelTaggingSettings(audioDescriptions[i].AudioChannelTaggingSettings),
			"audio_normalization_settings":   flattenMediaConvertAudioNormalizationSettings(audioDescriptions[i].AudioNormalizationSettings),
			"codec_settings":                 flattenMediaConvertAudioCodecSettings(audioDescriptions[i].CodecSettings),
			"remix_settings":                 flattenMediaConvertRemixSettings(audioDescriptions[i].RemixSettings),
		}
		audioDesc = append(audioDesc, m)
	}
	return audioDesc
}

func flattenMediaConvertAudioCodecSettings(cfg *mediaconvert.AudioCodecSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"codec":               aws.StringValue(cfg.Codec),
		"aac_settings":        flattenMediaConvertAacSettings(cfg.AacSettings),
		"ac3_settings":        flattenMediaConvertAc3Settings(cfg.Ac3Settings),
		"aiff_settings":       flattenMediaConvertAiffSettings(cfg.AiffSettings),
		"eac3_atmos_settings": flattenMediaConvertEac3AtmosSettings(cfg.Eac3AtmosSettings),
		"eac3_settings":       flattenMediaConvertEac3Settings(cfg.Eac3Settings),
		"mp2_settings":        flattenMediaConvertMp2Settings(cfg.Mp2Settings),
		"mp3_settings":        flattenMediaConvertMp3Settings(cfg.Mp3Settings),
		"opus_settings":       flattenMediaConvertOpusSettings(cfg.OpusSettings),
		"vorbis_settings":     flattenMediaConvertVorbisSettings(cfg.VorbisSettings),
		"wav_settings":        flattenMediaConvertWavSettings(cfg.WavSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertAudioNormalizationSettings(cfg *mediaconvert.AudioNormalizationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"algorithm":             aws.StringValue(cfg.Algorithm),
		"algorithm_control":     aws.StringValue(cfg.AlgorithmControl),
		"correction_gate_level": aws.Int64Value(cfg.CorrectionGateLevel),
		"loudness_logging":      aws.StringValue(cfg.LoudnessLogging),
		"peak_calculation":      aws.StringValue(cfg.PeakCalculation),
		"target_lkfs":           aws.Float64Value(cfg.TargetLkfs),
	}
	return []interface{}{m}
}

func flattenMediaConvertAacSettings(cfg *mediaconvert.AacSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_description_broadcaster_mix": aws.StringValue(cfg.AudioDescriptionBroadcasterMix),
		"bitrate":                           aws.Int64Value(cfg.Bitrate),
		"codec_profile":                     aws.StringValue(cfg.CodecProfile),
		"coding_mode":                       aws.StringValue(cfg.CodingMode),
		"rate_control_mode":                 aws.StringValue(cfg.RateControlMode),
		"raw_format":                        aws.StringValue(cfg.RawFormat),
		"sample_rate":                       aws.Int64Value(cfg.SampleRate),
		"specification":                     aws.StringValue(cfg.Specification),
		"vbr_quality":                       aws.StringValue(cfg.VbrQuality),
	}
	return []interface{}{m}
}

func flattenMediaConvertAc3Settings(cfg *mediaconvert.Ac3Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":                           aws.Int64Value(cfg.Bitrate),
		"bitstream_mode":                    aws.StringValue(cfg.BitstreamMode),
		"coding_mode":                       aws.StringValue(cfg.CodingMode),
		"dialnorm":                          aws.Int64Value(cfg.Dialnorm),
		"dynamic_range_compression_profile": aws.StringValue(cfg.DynamicRangeCompressionProfile),
		"lfe_filter":                        aws.StringValue(cfg.LfeFilter),
		"metadata_control":                  aws.StringValue(cfg.MetadataControl),
		"sample_rate":                       aws.Int64Value(cfg.SampleRate),
	}
	return []interface{}{m}
}

func flattenMediaConvertAiffSettings(cfg *mediaconvert.AiffSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitdepth":    aws.Int64Value(cfg.BitDepth),
		"channels":    aws.Int64Value(cfg.Channels),
		"sample_rate": aws.Int64Value(cfg.SampleRate),
	}
	return []interface{}{m}
}

func flattenMediaConvertEac3AtmosSettings(cfg *mediaconvert.Eac3AtmosSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":                        aws.Int64Value(cfg.Bitrate),
		"bitstream_mode":                 aws.StringValue(cfg.BitstreamMode),
		"coding_mode":                    aws.StringValue(cfg.CodingMode),
		"dialogue_intelligence":          aws.StringValue(cfg.DialogueIntelligence),
		"dynamic_range_compression_line": aws.StringValue(cfg.DynamicRangeCompressionLine),
		"dynamic_range_compression_rf":   aws.StringValue(cfg.DynamicRangeCompressionRf),
		"lo_ro_center_mix_level":         aws.Float64Value(cfg.LoRoCenterMixLevel),
		"lo_ro_surround_mix_level":       aws.Float64Value(cfg.LoRoSurroundMixLevel),
		"lt_rt_center_mix_level":         aws.Float64Value(cfg.LtRtCenterMixLevel),
		"lt_rt_surround_mix_level":       aws.Float64Value(cfg.LtRtSurroundMixLevel),
		"metering_mode":                  aws.StringValue(cfg.MeteringMode),
		"sample_rate":                    aws.Int64Value(cfg.SampleRate),
		"speech_threshold":               aws.Int64Value(cfg.SpeechThreshold),
		"stereo_downmix":                 aws.StringValue(cfg.StereoDownmix),
		"surround_ex_mode":               aws.StringValue(cfg.SurroundExMode),
	}
	return []interface{}{m}
}

func flattenMediaConvertEac3Settings(cfg *mediaconvert.Eac3Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"attenuation_control":            aws.StringValue(cfg.AttenuationControl),
		"bitrate":                        aws.Int64Value(cfg.Bitrate),
		"bitstream_mode":                 aws.StringValue(cfg.BitstreamMode),
		"coding_mode":                    aws.StringValue(cfg.CodingMode),
		"dc_filter":                      aws.StringValue(cfg.DcFilter),
		"dialnorm":                       aws.Int64Value(cfg.Dialnorm),
		"dynamic_range_compression_line": aws.StringValue(cfg.DynamicRangeCompressionLine),
		"dynamic_range_compression_rf":   aws.StringValue(cfg.DynamicRangeCompressionRf),
		"lfe_control":                    aws.StringValue(cfg.LfeControl),
		"lfe_filter":                     aws.StringValue(cfg.LfeFilter),
		"lo_ro_center_mix_level":         aws.Float64Value(cfg.LoRoCenterMixLevel),
		"lo_ro_surround_mix_level":       aws.Float64Value(cfg.LoRoSurroundMixLevel),
		"lt_rt_center_mix_level":         aws.Float64Value(cfg.LtRtCenterMixLevel),
		"lt_rt_surround_mix_level":       aws.Float64Value(cfg.LtRtSurroundMixLevel),
		"metadata_control":               aws.StringValue(cfg.MetadataControl),
		"passthrough_control":            aws.StringValue(cfg.PassthroughControl),
		"phase_control":                  aws.StringValue(cfg.PhaseControl),
		"sample_rate":                    aws.Int64Value(cfg.SampleRate),
		"stereo_downmix":                 aws.StringValue(cfg.StereoDownmix),
		"surround_ex_mode":               aws.StringValue(cfg.SurroundExMode),
		"surround_mode":                  aws.StringValue(cfg.SurroundMode),
	}
	return []interface{}{m}
}

func flattenMediaConvertMp2Settings(cfg *mediaconvert.Mp2Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":     aws.Int64Value(cfg.Bitrate),
		"channels":    aws.Int64Value(cfg.Channels),
		"sample_rate": aws.Int64Value(cfg.SampleRate),
	}
	return []interface{}{m}
}

func flattenMediaConvertMp3Settings(cfg *mediaconvert.Mp3Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":           aws.Int64Value(cfg.Bitrate),
		"channels":          aws.Int64Value(cfg.Channels),
		"rate_control_mode": aws.StringValue(cfg.RateControlMode),
		"sample_rate":       aws.Int64Value(cfg.SampleRate),
		"vbr_quality":       aws.Int64Value(cfg.VbrQuality),
	}
	return []interface{}{m}
}

func flattenMediaConvertOpusSettings(cfg *mediaconvert.OpusSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":     aws.Int64Value(cfg.Bitrate),
		"channels":    aws.Int64Value(cfg.Channels),
		"sample_rate": aws.Int64Value(cfg.SampleRate),
	}
	return []interface{}{m}
}

func flattenMediaConvertVorbisSettings(cfg *mediaconvert.VorbisSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"channels":    aws.Int64Value(cfg.Channels),
		"sample_rate": aws.Int64Value(cfg.SampleRate),
		"vbr_quality": aws.Int64Value(cfg.VbrQuality),
	}
	return []interface{}{m}
}

func flattenMediaConvertWavSettings(cfg *mediaconvert.WavSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitdepth":    aws.Int64Value(cfg.BitDepth),
		"channels":    aws.Int64Value(cfg.Channels),
		"format":      aws.StringValue(cfg.Format),
		"sample_rate": aws.Int64Value(cfg.SampleRate),
	}
	return []interface{}{m}
}

func flattenMediaConvertRemixSettings(cfg *mediaconvert.RemixSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"channel_mapping": map[string]interface{}{
			"output_channels": flattenMediaConvertOutputChannelMapping(cfg.ChannelMapping.OutputChannels),
		},
		"channels_in":  aws.Int64Value(cfg.ChannelsIn),
		"channels_out": aws.Int64Value(cfg.ChannelsOut),
	}
	return []interface{}{m}
}

func flattenMediaConvertAudioChannelTaggingSettings(cfg *mediaconvert.AudioChannelTaggingSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"channel_tag": aws.StringValue(cfg.ChannelTag),
	}
	return []interface{}{m}
}

func flattenMediaConvertOutputChannelMapping(cfg []*mediaconvert.OutputChannelMapping) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	ocResults := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		if cfg[i].InputChannels != nil {
			m := map[string]interface{}{
				"output_channels": map[string]interface{}{
					"input_channels": flattenInt64Set(cfg[i].InputChannels),
				},
			}
			ocResults = append(ocResults, m)
		}
	}
	return ocResults
}

func flattenMediaConvertCaptionDescriptionPreset(captionDescriptions []*mediaconvert.CaptionDescriptionPreset) []interface{} {
	if captionDescriptions == nil {
		return []interface{}{}
	}
	captionDescs := make([]interface{}, 0, len(captionDescriptions))
	for i := 0; i < len(captionDescriptions); i++ {
		m := map[string]interface{}{
			"custom_language_code": aws.StringValue(captionDescriptions[i].CustomLanguageCode),
			"destination_settings": flattenMediaConvertCaptionDestinationSettings(captionDescriptions[i].DestinationSettings),
			"language_code":        aws.StringValue(captionDescriptions[i].LanguageCode),
			"language_description": aws.StringValue(captionDescriptions[i].LanguageDescription),
		}
		captionDescs = append(captionDescs, m)
	}

	return captionDescs
}
func flattenMediaConvertCaptionDestinationSettings(cfg *mediaconvert.CaptionDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"destination_type":              aws.StringValue(cfg.DestinationType),
		"burnin_destination_settings":   flattenMediaConvertBurninDestinationSettings(cfg.BurninDestinationSettings),
		"dvb_sub_destination_settings":  flattenMediaConvertDvbSubDestinationSettings(cfg.DvbSubDestinationSettings),
		"embedded_destination_settings": flattenMediaConvertEmbeddedDestinationSettings(cfg.EmbeddedDestinationSettings),
		"imsc_destination_settings":     flattenMediaConvertImscDestinationSettings(cfg.ImscDestinationSettings),
		"scc_destination_settings":      flattenMediaConvertSccDestinationSettings(cfg.SccDestinationSettings),
		"teletext_destination_settings": flattenMediaConvertTeletextDestinationSettings(cfg.TeletextDestinationSettings),
		"ttml_destination_settings":     flattenMediaConvertTtmlDestinationSettings(cfg.TtmlDestinationSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertBurninDestinationSettings(cfg *mediaconvert.BurninDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"alignment":          aws.StringValue(cfg.Alignment),
		"background_color":   aws.StringValue(cfg.BackgroundColor),
		"background_opacity": aws.Int64Value(cfg.BackgroundOpacity),
		"font_color":         aws.StringValue(cfg.FontColor),
		"font_opacity":       aws.Int64Value(cfg.FontOpacity),
		"font_resolution":    aws.Int64Value(cfg.FontResolution),
		"font_script":        aws.StringValue(cfg.FontColor),
		"font_size":          aws.Int64Value(cfg.FontSize),
		"outline_color":      aws.StringValue(cfg.OutlineColor),
		"outline_size":       aws.Int64Value(cfg.OutlineSize),
		"shadow_color":       aws.StringValue(cfg.ShadowColor),
		"shadow_opacity":     aws.Int64Value(cfg.ShadowOpacity),
		"shadow_x_offset":    aws.Int64Value(cfg.ShadowXOffset),
		"shadow_y_offset":    aws.Int64Value(cfg.ShadowYOffset),
		"teletext_spacing":   aws.StringValue(cfg.TeletextSpacing),
		"x_position":         aws.Int64Value(cfg.XPosition),
		"y_position":         aws.Int64Value(cfg.YPosition),
	}
	return []interface{}{m}
}

func flattenMediaConvertDvbSubDestinationSettings(cfg *mediaconvert.DvbSubDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"alignment":          aws.StringValue(cfg.Alignment),
		"background_color":   aws.StringValue(cfg.BackgroundColor),
		"background_opacity": aws.Int64Value(cfg.BackgroundOpacity),
		"font_color":         aws.StringValue(cfg.FontColor),
		"font_opacity":       aws.Int64Value(cfg.FontOpacity),
		"font_resolution":    aws.Int64Value(cfg.FontResolution),
		"font_script":        aws.StringValue(cfg.FontColor),
		"font_size":          aws.Int64Value(cfg.FontSize),
		"outline_color":      aws.StringValue(cfg.OutlineColor),
		"outline_size":       aws.Int64Value(cfg.OutlineSize),
		"shadow_color":       aws.StringValue(cfg.ShadowColor),
		"shadow_opacity":     aws.Int64Value(cfg.ShadowOpacity),
		"shadow_x_offset":    aws.Int64Value(cfg.ShadowXOffset),
		"shadow_y_offset":    aws.Int64Value(cfg.ShadowYOffset),
		"subtitling_type":    aws.StringValue(cfg.SubtitlingType),
		"teletext_spacing":   aws.StringValue(cfg.TeletextSpacing),
		"x_position":         aws.Int64Value(cfg.XPosition),
		"y_position":         aws.Int64Value(cfg.YPosition),
	}
	return []interface{}{m}
}

func flattenMediaConvertEmbeddedDestinationSettings(cfg *mediaconvert.EmbeddedDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"destination_608_channel_number": aws.Int64Value(cfg.Destination608ChannelNumber),
		"destination_708_service_number": aws.Int64Value(cfg.Destination708ServiceNumber),
	}
	return []interface{}{m}
}

func flattenMediaConvertImscDestinationSettings(cfg *mediaconvert.ImscDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"style_passthrough": aws.StringValue(cfg.StylePassthrough),
	}
	return []interface{}{m}
}

func flattenMediaConvertSccDestinationSettings(cfg *mediaconvert.SccDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate": aws.StringValue(cfg.Framerate),
	}
	return []interface{}{m}
}

func flattenMediaConvertTeletextDestinationSettings(cfg *mediaconvert.TeletextDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"page_number": aws.StringValue(cfg.PageNumber),
		"page_types":  flattenStringSet(cfg.PageTypes),
	}
	return []interface{}{m}
}

func flattenMediaConvertTtmlDestinationSettings(cfg *mediaconvert.TtmlDestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"style_passthrough": aws.StringValue(cfg.StylePassthrough),
	}
	return []interface{}{m}
}

func flattenMediaConvertContainerSettings(containerSettings *mediaconvert.ContainerSettings) []interface{} {
	if containerSettings == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"cmfc_settings": flattenMediaConvertCmfcSettings(containerSettings.CmfcSettings),
		"container":     aws.StringValue(containerSettings.Container),
		"f4v_settings":  flattenMediaConvertF4vSettings(containerSettings.F4vSettings),
		"m2ts_settings": flattenMediaConvertM2tsSettings(containerSettings.M2tsSettings),
		"m3u8_settings": flattenMediaConvertM3u8Settings(containerSettings.M3u8Settings),
		"mov_settings":  flattenMediaConvertMovSettings(containerSettings.MovSettings),
		"mp4_settings":  flattenMediaConvertMp4Settings(containerSettings.Mp4Settings),
		"mpd_settings":  flattenMediaConvertMpdSettings(containerSettings.MpdSettings),
		"mxf_settings":  flattenMediaConvertMxfSettings(containerSettings.MxfSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertCmfcSettings(cfg *mediaconvert.CmfcSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_duration": aws.StringValue(cfg.AudioDuration),
		"scte35_esam":    aws.StringValue(cfg.Scte35Esam),
		"scte35_source":  aws.StringValue(cfg.Scte35Source),
	}
	return []interface{}{m}
}

func flattenMediaConvertF4vSettings(cfg *mediaconvert.F4vSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"moov_placement": aws.StringValue(cfg.MoovPlacement),
	}
	return []interface{}{m}
}

func flattenMediaConvertM2tsSettings(cfg *mediaconvert.M2tsSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_buffer_model":       aws.StringValue(cfg.AudioBufferModel),
		"audio_duration":           aws.StringValue(cfg.AudioDuration),
		"audio_frames_per_pes":     aws.Int64Value(cfg.AudioFramesPerPes),
		"audio_pids":               flattenInt64Set(cfg.AudioPids),
		"bitrate":                  aws.Int64Value(cfg.Bitrate),
		"buffer_model":             aws.StringValue(cfg.BufferModel),
		"dvb_nit_settings":         flattenMediaConvertDvbNitSettings(cfg.DvbNitSettings),
		"dvb_sdt_settings":         flattenMediaConvertDvbSdtSettings(cfg.DvbSdtSettings),
		"dvb_sub_pids":             flattenInt64Set(cfg.DvbSubPids),
		"dvb_tdt_settings":         flattenMediaConvertDvbTdtSettings(cfg.DvbTdtSettings),
		"dvb_teletext_pid":         aws.Int64Value(cfg.DvbTeletextPid),
		"ebp_audio_interval":       aws.StringValue(cfg.EbpAudioInterval),
		"ebp_placement":            aws.StringValue(cfg.EbpPlacement),
		"es_rate_in_pes":           aws.StringValue(cfg.EsRateInPes),
		"force_ts_video_ebp_order": aws.StringValue(cfg.ForceTsVideoEbpOrder),
		"fragment_time":            aws.Float64Value(cfg.FragmentTime),
		"max_pcr_interval":         aws.Int64Value(cfg.MaxPcrInterval),
		"min_ebp_interval":         aws.Int64Value(cfg.MinEbpInterval),
		"nielsen_id3":              aws.StringValue(cfg.NielsenId3),
		"null_packet_bitrate":      aws.Float64Value(cfg.NullPacketBitrate),
		"pat_interval":             aws.Int64Value(cfg.PatInterval),
		"pcr_control":              aws.StringValue(cfg.PcrControl),
		"pcr_pid":                  aws.Int64Value(cfg.PcrPid),
		"pmt_interval":             aws.Int64Value(cfg.PmtInterval),
		"pmt_pid":                  aws.Int64Value(cfg.PmtPid),
		"private_metadata_pid":     aws.Int64Value(cfg.PrivateMetadataPid),
		"program_number":           aws.Int64Value(cfg.ProgramNumber),
		"rate_mode":                aws.StringValue(cfg.RateMode),
		"scte_35_esam":             flattenMediaConvertM2tsScte35Esam(cfg.Scte35Esam),
		"scte_35_pid":              aws.Int64Value(cfg.Scte35Pid),
		"scte_35_source":           aws.StringValue(cfg.Scte35Source),
		"segmentation_markers":     aws.StringValue(cfg.SegmentationMarkers),
		"segmentation_style":       aws.StringValue(cfg.SegmentationStyle),
		"segmentation_time":        aws.Float64Value(cfg.SegmentationTime),
		"timed_metadata_pid":       aws.Int64Value(cfg.TimedMetadataPid),
		"transport_stream_id":      aws.Int64Value(cfg.TransportStreamId),
		"video_pid":                aws.Int64Value(cfg.VideoPid),
	}
	return []interface{}{m}
}

func flattenMediaConvertDvbNitSettings(cfg *mediaconvert.DvbNitSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"network_id":   aws.Int64Value(cfg.NetworkId),
		"network_name": aws.StringValue(cfg.NetworkName),
		"nit_interval": aws.Int64Value(cfg.NitInterval),
	}
	return []interface{}{m}
}

func flattenMediaConvertDvbSdtSettings(cfg *mediaconvert.DvbSdtSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"output_sdt":            aws.StringValue(cfg.OutputSdt),
		"sdt_interval":          aws.Int64Value(cfg.SdtInterval),
		"service_name":          aws.StringValue(cfg.ServiceName),
		"service_provider_name": aws.StringValue(cfg.ServiceProviderName),
	}
	return []interface{}{m}
}

func flattenMediaConvertDvbTdtSettings(cfg *mediaconvert.DvbTdtSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"tdt_interval": aws.Int64Value(cfg.TdtInterval),
	}
	return []interface{}{m}
}

func flattenMediaConvertM2tsScte35Esam(cfg *mediaconvert.M2tsScte35Esam) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"scte_35_esam_pid": aws.Int64Value(cfg.Scte35EsamPid),
	}
	return []interface{}{m}
}

func flattenMediaConvertM3u8Settings(cfg *mediaconvert.M3u8Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_duration":       aws.StringValue(cfg.AudioDuration),
		"audio_frames_per_pes": aws.Int64Value(cfg.AudioFramesPerPes),
		"audio_pids":           flattenInt64Set(cfg.AudioPids),
		"nielsen_id3":          aws.StringValue(cfg.NielsenId3),
		"pat_interval":         aws.Int64Value(cfg.PatInterval),
		"pcr_control":          aws.StringValue(cfg.PcrControl),
		"pcr_pid":              aws.Int64Value(cfg.PcrPid),
		"pmt_interval":         aws.Int64Value(cfg.PmtInterval),
		"pmt_pid":              aws.Int64Value(cfg.PmtPid),
		"private_metadata_pid": aws.Int64Value(cfg.PrivateMetadataPid),
		"program_number":       aws.Int64Value(cfg.ProgramNumber),
		"scte_35_pid":          aws.Int64Value(cfg.Scte35Pid),
		"scte_35_source":       aws.StringValue(cfg.Scte35Source),
		"timed_metadata":       aws.StringValue(cfg.TimedMetadata),
		"timed_metadata_pid":   aws.Int64Value(cfg.TimedMetadataPid),
		"transport_stream_id":  aws.Int64Value(cfg.TransportStreamId),
		"video_pid":            aws.Int64Value(cfg.VideoPid),
	}
	return []interface{}{m}
}

func flattenMediaConvertMovSettings(cfg *mediaconvert.MovSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"clap_atom":            aws.StringValue(cfg.ClapAtom),
		"cslg_atom":            aws.StringValue(cfg.CslgAtom),
		"mpeg2_fourcc_control": aws.StringValue(cfg.Mpeg2FourCCControl),
		"padding_control":      aws.StringValue(cfg.PaddingControl),
		"reference":            aws.StringValue(cfg.Reference),
	}
	return []interface{}{m}
}

func flattenMediaConvertMp4Settings(cfg *mediaconvert.Mp4Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_duration":  aws.StringValue(cfg.AudioDuration),
		"cslg_atom":       aws.StringValue(cfg.CslgAtom),
		"ctts_version":    aws.Int64Value(cfg.CttsVersion),
		"free_space_box":  aws.StringValue(cfg.FreeSpaceBox),
		"moov_placement":  aws.StringValue(cfg.MoovPlacement),
		"mp4_major_brand": aws.StringValue(cfg.Mp4MajorBrand),
	}
	return []interface{}{m}
}

func flattenMediaConvertMpdSettings(cfg *mediaconvert.MpdSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"accessibility_caption_hints": aws.StringValue(cfg.AccessibilityCaptionHints),
		"audio_duration":              aws.StringValue(cfg.AudioDuration),
		"caption_container_type":      aws.StringValue(cfg.CaptionContainerType),
		"scte_35_esam":                aws.StringValue(cfg.Scte35Esam),
		"scte_35_source":              aws.StringValue(cfg.Scte35Source),
	}
	return []interface{}{m}
}

func flattenMediaConvertMxfSettings(cfg *mediaconvert.MxfSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"afd_signaling": aws.StringValue(cfg.AfdSignaling),
		"profile":       aws.StringValue(cfg.Profile),
	}
	return []interface{}{m}
}

func flattenMediaConvertVideoDescription(videoDescription *mediaconvert.VideoDescription) []interface{} {
	if videoDescription == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"afd_signaling":       aws.StringValue(videoDescription.AfdSignaling),
		"anti_alias":          aws.StringValue(videoDescription.AntiAlias),
		"codec_settings":      flattenMediaConvertVideoCodecSettings(videoDescription.CodecSettings),
		"color_metadata":      aws.StringValue(videoDescription.ColorMetadata),
		"crop":                flattenMediaConvertRectangle(videoDescription.Crop),
		"drop_frame_timecode": aws.StringValue(videoDescription.DropFrameTimecode),
		"fixed_afd":           aws.Int64Value(videoDescription.FixedAfd),
		"height":              aws.Int64Value(videoDescription.Height),
		"position":            flattenMediaConvertRectangle(videoDescription.Position),
		"respond_to_afd":      aws.StringValue(videoDescription.RespondToAfd),
		"scaling_behavior":    aws.StringValue(videoDescription.ScalingBehavior),
		"sharpness":           aws.Int64Value(videoDescription.Sharpness),
		"timecode_insertion":  aws.StringValue(videoDescription.TimecodeInsertion),
		"video_preprocessors": flattenMediaConvertVideoPreprocessor(videoDescription.VideoPreprocessors),
		"width":               aws.Int64Value(videoDescription.Width),
	}
	return []interface{}{m}
}

func flattenMediaConvertVideoCodecSettings(cfg *mediaconvert.VideoCodecSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"av1_settings":           flattenMediaConvertAv1Settings(cfg.Av1Settings),
		"avc_intra_settings":     flattenMediaConvertAvcIntraSettings(cfg.AvcIntraSettings),
		"codec":                  aws.StringValue(cfg.Codec),
		"frame_capture_settings": flattenMediaConvertFrameCaptureSettings(cfg.FrameCaptureSettings),
		"h264_settings":          flattenMediaConvertH264Settings(cfg.H264Settings),
		"h265_settings":          flattenMediaConvertH265Settings(cfg.H265Settings),
		"mpeg2_settings":         flattenMediaConvertMpeg2Settings(cfg.Mpeg2Settings),
		"prores_settings":        flattenMediaConvertProresSettings(cfg.ProresSettings),
		"vc3_settings":           flattenMediaConvertVc3Settings(cfg.Vc3Settings),
		"vp8_settings":           flattenMediaConvertVp8Settings(cfg.Vp8Settings),
		"vp9_settings":           flattenMediaConvertVp9Settings(cfg.Vp9Settings),
	}
	return []interface{}{m}
}

func flattenMediaConvertAv1Settings(cfg *mediaconvert.Av1Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"adaptive_quantization":                    aws.StringValue(cfg.AdaptiveQuantization),
		"framerate_control":                        aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm":           aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":                    aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":                      aws.Int64Value(cfg.FramerateNumerator),
		"gop_size":                                 aws.Float64Value(cfg.GopSize),
		"max_bitrate":                              aws.Int64Value(cfg.MaxBitrate),
		"number_b_frames_between_reference_frames": aws.Int64Value(cfg.NumberBFramesBetweenReferenceFrames),
		"qvbr_settings":                            flattenMediaConvertAv1QvbrSettings(cfg.QvbrSettings),
		"rate_control_mode":                        aws.StringValue(cfg.RateControlMode),
		"slices":                                   aws.Int64Value(cfg.Slices),
		"spatial_adaptive_quantization":            aws.StringValue(cfg.SpatialAdaptiveQuantization),
	}
	return []interface{}{m}
}

func flattenMediaConvertAv1QvbrSettings(cfg *mediaconvert.Av1QvbrSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"qvbr_quality_level":           aws.Int64Value(cfg.QvbrQualityLevel),
		"qvbr_quality_level_fine_tune": aws.Float64Value(cfg.QvbrQualityLevelFineTune),
	}
	return []interface{}{m}
}

func flattenMediaConvertAvcIntraSettings(cfg *mediaconvert.AvcIntraSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"avc_intra_class":                aws.StringValue(cfg.AvcIntraClass),
		"framerate_control":              aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm": aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":          aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":            aws.Int64Value(cfg.FramerateNumerator),
		"interlace_mode":                 aws.StringValue(cfg.InterlaceMode),
		"slow_pal":                       aws.StringValue(cfg.SlowPal),
		"telecine":                       aws.StringValue(cfg.Telecine),
	}
	return []interface{}{m}
}

func flattenMediaConvertFrameCaptureSettings(cfg *mediaconvert.FrameCaptureSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate_denominator": aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":   aws.Int64Value(cfg.FramerateNumerator),
		"max_captures":          aws.Int64Value(cfg.MaxCaptures),
		"quality":               aws.Int64Value(cfg.Quality),
	}
	return []interface{}{m}
}

func flattenMediaConvertH264Settings(cfg *mediaconvert.H264Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"adaptive_quantization":                    aws.StringValue(cfg.AdaptiveQuantization),
		"bitrate":                                  aws.Int64Value(cfg.Bitrate),
		"codec_level":                              aws.StringValue(cfg.CodecLevel),
		"codec_profile":                            aws.StringValue(cfg.CodecProfile),
		"dynamic_sub_gop":                          aws.StringValue(cfg.DynamicSubGop),
		"entropy_encoding":                         aws.StringValue(cfg.EntropyEncoding),
		"field_encoding":                           aws.StringValue(cfg.FieldEncoding),
		"flicker_adaptive_quantization":            aws.StringValue(cfg.FlickerAdaptiveQuantization),
		"framerate_control":                        aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm":           aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":                    aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":                      aws.Int64Value(cfg.FramerateNumerator),
		"gop_b_reference":                          aws.StringValue(cfg.GopBReference),
		"gop_closed_cadence":                       aws.Int64Value(cfg.GopClosedCadence),
		"gop_size":                                 aws.Float64Value(cfg.GopSize),
		"gop_size_units":                           aws.StringValue(cfg.GopSizeUnits),
		"hrd_buffer_initial_fill_percentage":       aws.Int64Value(cfg.HrdBufferInitialFillPercentage),
		"hrd_buffer_size":                          aws.Int64Value(cfg.HrdBufferSize),
		"interlace_mode":                           aws.StringValue(cfg.InterlaceMode),
		"max_bitrate":                              aws.Int64Value(cfg.MaxBitrate),
		"min_i_interval":                           aws.Int64Value(cfg.MinIInterval),
		"number_b_frames_between_reference_frames": aws.Int64Value(cfg.NumberBFramesBetweenReferenceFrames),
		"number_reference_frames":                  aws.Int64Value(cfg.NumberReferenceFrames),
		"par_control":                              aws.StringValue(cfg.ParControl),
		"par_denominator":                          aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                            aws.Int64Value(cfg.ParNumerator),
		"quality_tuning_level":                     aws.StringValue(cfg.QualityTuningLevel),
		"qvbr_settings":                            flattenMediaConvertH264QvbrSettings(cfg.QvbrSettings),
		"rate_control_mode":                        aws.StringValue(cfg.RateControlMode),
		"repeat_pps":                               aws.StringValue(cfg.RepeatPps),
		"scene_change_detect":                      aws.StringValue(cfg.SceneChangeDetect),
		"slices":                                   aws.Int64Value(cfg.Slices),
		"slow_pal":                                 aws.StringValue(cfg.SlowPal),
		"softness":                                 aws.Int64Value(cfg.Softness),
		"spatial_adaptive_quantization":            aws.StringValue(cfg.SpatialAdaptiveQuantization),
		"syntax":                                   aws.StringValue(cfg.Syntax),
		"telecine":                                 aws.StringValue(cfg.Telecine),
		"temporal_adaptive_quantization":           aws.StringValue(cfg.TemporalAdaptiveQuantization),
		"unregistered_sei_timecode":                aws.StringValue(cfg.UnregisteredSeiTimecode),
	}
	return []interface{}{m}
}
func flattenMediaConvertH264QvbrSettings(cfg *mediaconvert.H264QvbrSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"max_average_bitrate":          aws.Int64Value(cfg.MaxAverageBitrate),
		"qvbr_quality_level":           aws.Int64Value(cfg.QvbrQualityLevel),
		"qvbr_quality_level_fine_tune": aws.Float64Value(cfg.QvbrQualityLevelFineTune),
	}
	return []interface{}{m}
}

func flattenMediaConvertH265Settings(cfg *mediaconvert.H265Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"adaptive_quantization":                    aws.StringValue(cfg.AdaptiveQuantization),
		"alternate_transfer_function_sei":          aws.StringValue(cfg.AlternateTransferFunctionSei),
		"bitrate":                                  aws.Int64Value(cfg.Bitrate),
		"codec_level":                              aws.StringValue(cfg.CodecLevel),
		"codec_profile":                            aws.StringValue(cfg.CodecProfile),
		"dynamic_sub_gop":                          aws.StringValue(cfg.DynamicSubGop),
		"flicker_adaptive_quantization":            aws.StringValue(cfg.FlickerAdaptiveQuantization),
		"framerate_control":                        aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm":           aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":                    aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":                      aws.Int64Value(cfg.FramerateNumerator),
		"gop_b_reference":                          aws.StringValue(cfg.GopBReference),
		"gop_closed_cadence":                       aws.Int64Value(cfg.GopClosedCadence),
		"gop_size":                                 aws.Float64Value(cfg.GopSize),
		"gop_size_units":                           aws.StringValue(cfg.GopSizeUnits),
		"hrd_buffer_initial_fill_percentage":       aws.Int64Value(cfg.HrdBufferInitialFillPercentage),
		"hrd_buffer_size":                          aws.Int64Value(cfg.HrdBufferSize),
		"interlace_mode":                           aws.StringValue(cfg.InterlaceMode),
		"max_bitrate":                              aws.Int64Value(cfg.MaxBitrate),
		"min_i_interval":                           aws.Int64Value(cfg.MinIInterval),
		"number_b_frames_between_reference_frames": aws.Int64Value(cfg.NumberBFramesBetweenReferenceFrames),
		"number_reference_frames":                  aws.Int64Value(cfg.NumberReferenceFrames),
		"par_control":                              aws.StringValue(cfg.ParControl),
		"par_denominator":                          aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                            aws.Int64Value(cfg.ParNumerator),
		"quality_tuning_level":                     aws.StringValue(cfg.QualityTuningLevel),
		"qvbr_settings":                            flattenMediaConvertH265QvbrSettings(cfg.QvbrSettings),
		"rate_control_mode":                        aws.StringValue(cfg.RateControlMode),
		"sample_adaptive_offset_filter_mode":       aws.StringValue(cfg.SampleAdaptiveOffsetFilterMode),
		"scene_change_detect":                      aws.StringValue(cfg.SceneChangeDetect),
		"slices":                                   aws.Int64Value(cfg.Slices),
		"slow_pal":                                 aws.StringValue(cfg.SlowPal),
		"spatial_adaptive_quantization":            aws.StringValue(cfg.SpatialAdaptiveQuantization),
		"telecine":                                 aws.StringValue(cfg.Telecine),
		"temporal_adaptive_quantization":           aws.StringValue(cfg.TemporalAdaptiveQuantization),
		"temporal_ids":                             aws.StringValue(cfg.TemporalIds),
		"tiles":                                    aws.StringValue(cfg.Tiles),
		"unregistered_sei_timecode":                aws.StringValue(cfg.UnregisteredSeiTimecode),
		"write_mp4_packaging_type":                 aws.StringValue(cfg.WriteMp4PackagingType),
	}
	return []interface{}{m}
}
func flattenMediaConvertH265QvbrSettings(cfg *mediaconvert.H265QvbrSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"max_average_bitrate":          aws.Int64Value(cfg.MaxAverageBitrate),
		"qvbr_quality_level":           aws.Int64Value(cfg.QvbrQualityLevel),
		"qvbr_quality_level_fine_tune": aws.Float64Value(cfg.QvbrQualityLevelFineTune),
	}
	return []interface{}{m}
}

func flattenMediaConvertMpeg2Settings(cfg *mediaconvert.Mpeg2Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"adaptive_quantization":                    aws.StringValue(cfg.AdaptiveQuantization),
		"bitrate":                                  aws.Int64Value(cfg.Bitrate),
		"codec_level":                              aws.StringValue(cfg.CodecLevel),
		"codec_profile":                            aws.StringValue(cfg.CodecProfile),
		"dynamic_sub_gop":                          aws.StringValue(cfg.DynamicSubGop),
		"framerate_control":                        aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm":           aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":                    aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":                      aws.Int64Value(cfg.FramerateNumerator),
		"gop_closed_cadence":                       aws.Int64Value(cfg.GopClosedCadence),
		"gop_size":                                 aws.Float64Value(cfg.GopSize),
		"gop_size_units":                           aws.StringValue(cfg.GopSizeUnits),
		"hrd_buffer_initial_fill_percentage":       aws.Int64Value(cfg.HrdBufferInitialFillPercentage),
		"hrd_buffer_size":                          aws.Int64Value(cfg.HrdBufferSize),
		"interlace_mode":                           aws.StringValue(cfg.InterlaceMode),
		"intra_dc_precision":                       aws.StringValue(cfg.IntraDcPrecision),
		"max_bitrate":                              aws.Int64Value(cfg.MaxBitrate),
		"min_i_interval":                           aws.Int64Value(cfg.MinIInterval),
		"number_b_frames_between_reference_frames": aws.Int64Value(cfg.NumberBFramesBetweenReferenceFrames),
		"par_control":                              aws.StringValue(cfg.ParControl),
		"par_denominator":                          aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                            aws.Int64Value(cfg.ParNumerator),
		"quality_tuning_level":                     aws.StringValue(cfg.QualityTuningLevel),
		"rate_control_mode":                        aws.StringValue(cfg.RateControlMode),
		"scene_change_detect":                      aws.StringValue(cfg.SceneChangeDetect),
		"slowpal":                                  aws.StringValue(cfg.SlowPal),
		"softness":                                 aws.Int64Value(cfg.Softness),
		"spatial_adaptive_quantization":            aws.StringValue(cfg.SpatialAdaptiveQuantization),
		"syntax":                                   aws.StringValue(cfg.Syntax),
		"telecine":                                 aws.StringValue(cfg.Telecine),
		"temporal_adaptive_quantization":           aws.StringValue(cfg.TemporalAdaptiveQuantization),
	}
	return []interface{}{m}
}

func flattenMediaConvertProresSettings(cfg *mediaconvert.ProresSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"codec_profile":                  aws.StringValue(cfg.CodecProfile),
		"framerate_control":              aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm": aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":          aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":            aws.Int64Value(cfg.FramerateNumerator),
		"interlace_mode":                 aws.StringValue(cfg.InterlaceMode),
		"par_control":                    aws.StringValue(cfg.ParControl),
		"par_denominator":                aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                  aws.Int64Value(cfg.ParNumerator),
		"slowpal":                        aws.StringValue(cfg.SlowPal),
		"telecine":                       aws.StringValue(cfg.Telecine),
	}
	return []interface{}{m}
}

func flattenMediaConvertVc3Settings(cfg *mediaconvert.Vc3Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate_control":              aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm": aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":          aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":            aws.Int64Value(cfg.FramerateNumerator),
		"interlace_mode":                 aws.StringValue(cfg.InterlaceMode),
		"slowpal":                        aws.StringValue(cfg.SlowPal),
		"telecine":                       aws.StringValue(cfg.Telecine),
		"vc3_class":                      aws.StringValue(cfg.Vc3Class),
	}
	return []interface{}{m}
}

func flattenMediaConvertVp8Settings(cfg *mediaconvert.Vp8Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":                        aws.Int64Value(cfg.Bitrate),
		"framerate_control":              aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm": aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":          aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":            aws.Int64Value(cfg.FramerateNumerator),
		"gop_size":                       aws.Float64Value(cfg.GopSize),
		"hrd_buffer_size":                aws.Int64Value(cfg.HrdBufferSize),
		"max_bitrate":                    aws.Int64Value(cfg.MaxBitrate),
		"par_control":                    aws.StringValue(cfg.ParControl),
		"par_denominator":                aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                  aws.Int64Value(cfg.ParNumerator),
		"quality_tuning_level":           aws.StringValue(cfg.QualityTuningLevel),
		"rate_control_mode":              aws.StringValue(cfg.RateControlMode),
	}
	return []interface{}{m}
}

func flattenMediaConvertVp9Settings(cfg *mediaconvert.Vp9Settings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"bitrate":                        aws.Int64Value(cfg.Bitrate),
		"framerate_control":              aws.StringValue(cfg.FramerateControl),
		"framerate_conversion_algorithm": aws.StringValue(cfg.FramerateConversionAlgorithm),
		"framerate_denominator":          aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":            aws.Int64Value(cfg.FramerateNumerator),
		"gop_size":                       aws.Float64Value(cfg.GopSize),
		"hrd_buffer_size":                aws.Int64Value(cfg.HrdBufferSize),
		"max_bitrate":                    aws.Int64Value(cfg.MaxBitrate),
		"par_control":                    aws.StringValue(cfg.ParControl),
		"par_denominator":                aws.Int64Value(cfg.ParDenominator),
		"par_numerator":                  aws.Int64Value(cfg.ParNumerator),
		"quality_tuning_level":           aws.StringValue(cfg.QualityTuningLevel),
		"rate_control_mode":              aws.StringValue(cfg.RateControlMode),
	}
	return []interface{}{m}
}

func flattenMediaConvertRectangle(cfg *mediaconvert.Rectangle) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"height": aws.Int64Value(cfg.Height),
		"width":  aws.Int64Value(cfg.Width),
		"x":      aws.Int64Value(cfg.X),
		"y":      aws.Int64Value(cfg.Y),
	}
	return []interface{}{m}
}

func flattenMediaConvertVideoPreprocessor(cfg *mediaconvert.VideoPreprocessor) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"color_corrector":     flattenMediaConvertColorCorrector(cfg.ColorCorrector),
		"deinterlacer":        flattenMediaConvertDeinterlacer(cfg.Deinterlacer),
		"dolby_vision":        flattenMediaConvertDolbyVision(cfg.DolbyVision),
		"image_inserter":      flattenMediaConvertImageInserter(cfg.ImageInserter),
		"noise_reducer":       flattenMediaConvertNoiseReducer(cfg.NoiseReducer),
		"partner_watermaking": flattenMediaConvertPartnerWatermarking(cfg.PartnerWatermarking),
		"timecode_burnin":     flattenMediaConvertTimecodeBurnin(cfg.TimecodeBurnin),
	}
	return []interface{}{m}
}

func flattenMediaConvertColorCorrector(cfg *mediaconvert.ColorCorrector) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"brightness":             aws.Int64Value(cfg.Brightness),
		"color_space_conversion": aws.StringValue(cfg.ColorSpaceConversion),
		"contrast":               aws.Int64Value(cfg.Contrast),
		"hdr10_metadata":         flattenMediaConvertHdr10Metadata(cfg.Hdr10Metadata),
		"hue":                    aws.Int64Value(cfg.Hue),
		"saturation":             aws.Int64Value(cfg.Saturation),
	}
	return []interface{}{m}
}

func flattenMediaConvertHdr10Metadata(cfg *mediaconvert.Hdr10Metadata) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"blue_primary_x":                aws.Int64Value(cfg.BluePrimaryX),
		"blue_primary_y":                aws.Int64Value(cfg.BluePrimaryY),
		"green_primary_x":               aws.Int64Value(cfg.GreenPrimaryX),
		"green_primary_y":               aws.Int64Value(cfg.GreenPrimaryY),
		"max_content_light_level":       aws.Int64Value(cfg.MaxContentLightLevel),
		"max_frame_average_light_level": aws.Int64Value(cfg.MaxFrameAverageLightLevel),
		"max_luminance":                 aws.Int64Value(cfg.MaxLuminance),
		"min_luminance":                 aws.Int64Value(cfg.MinLuminance),
		"red_primary_x":                 aws.Int64Value(cfg.RedPrimaryX),
		"red_primary_y":                 aws.Int64Value(cfg.RedPrimaryY),
		"white_point_x":                 aws.Int64Value(cfg.WhitePointX),
		"white_point_y":                 aws.Int64Value(cfg.WhitePointY),
	}
	return []interface{}{m}
}

func flattenMediaConvertDeinterlacer(cfg *mediaconvert.Deinterlacer) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"algorithm": aws.StringValue(cfg.Algorithm),
		"control":   aws.StringValue(cfg.Control),
		"mode":      aws.StringValue(cfg.Mode),
	}
	return []interface{}{m}
}

func flattenMediaConvertDolbyVision(cfg *mediaconvert.DolbyVision) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"l6_metadata": flattenMediaConvertDolbyVisionLevel6Metadata(cfg.L6Metadata),
		"l6_mode":     aws.StringValue(cfg.L6Mode),
		"profile":     aws.StringValue(cfg.Profile),
	}
	return []interface{}{m}
}

func flattenMediaConvertDolbyVisionLevel6Metadata(cfg *mediaconvert.DolbyVisionLevel6Metadata) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"max_cll":  aws.Int64Value(cfg.MaxCll),
		"max_fall": aws.Int64Value(cfg.MaxFall),
	}
	return []interface{}{m}
}

func flattenMediaConvertImageInserter(cfg *mediaconvert.ImageInserter) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"insertable_image": flattenMediaConvertInsertableImages(cfg.InsertableImages),
	}
	return []interface{}{m}
}

func flattenMediaConvertNoiseReducer(cfg *mediaconvert.NoiseReducer) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"filter":                   aws.StringValue(cfg.Filter),
		"filter_settings":          flattenMediaConvertNoiseReducerFilterSettings(cfg.FilterSettings),
		"spatial_filter_settings":  flattenMediaConvertNoiseReducerSpatialFilterSettings(cfg.SpatialFilterSettings),
		"temporal_filter_settings": flattenMediaConvertNoiseReducerTemporalFilterSettings(cfg.TemporalFilterSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertNoiseReducerFilterSettings(cfg *mediaconvert.NoiseReducerFilterSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"strength": aws.Int64Value(cfg.Strength),
	}
	return []interface{}{m}
}

func flattenMediaConvertNoiseReducerSpatialFilterSettings(cfg *mediaconvert.NoiseReducerSpatialFilterSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"post_filter_sharpen_strength": aws.Int64Value(cfg.PostFilterSharpenStrength),
		"speed":                        aws.Int64Value(cfg.Speed),
		"strength":                     aws.Int64Value(cfg.Strength),
	}
	return []interface{}{m}
}
func flattenMediaConvertNoiseReducerTemporalFilterSettings(cfg *mediaconvert.NoiseReducerTemporalFilterSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"aggressive_mode":          aws.Int64Value(cfg.AggressiveMode),
		"post_temporal_sharpening": aws.StringValue(cfg.PostTemporalSharpening),
		"speed":                    aws.Int64Value(cfg.Speed),
		"strength":                 aws.Int64Value(cfg.Strength),
	}
	return []interface{}{m}
}

func flattenMediaConvertPartnerWatermarking(cfg *mediaconvert.PartnerWatermarking) []interface{} {
	if cfg == nil && cfg.NexguardFileMarkerSettings != nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"nexguard_file_marker_settings": map[string]interface{}{
			"license":  aws.StringValue(cfg.NexguardFileMarkerSettings.License),
			"payload":  aws.Int64Value(cfg.NexguardFileMarkerSettings.Payload),
			"preset":   aws.StringValue(cfg.NexguardFileMarkerSettings.Preset),
			"strength": aws.StringValue(cfg.NexguardFileMarkerSettings.Strength),
		},
	}
	return []interface{}{m}
}

func flattenMediaConvertTimecodeBurnin(cfg *mediaconvert.TimecodeBurnin) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"font_size": aws.Int64Value(cfg.FontSize),
		"position":  aws.StringValue(cfg.Position),
		"prefix":    aws.StringValue(cfg.Prefix),
	}
	return []interface{}{m}
}

func flattenMediaConvertInsertableImages(insertableImages []*mediaconvert.InsertableImage) []interface{} {
	if insertableImages == nil {
		return []interface{}{}
	}
	insImgs := make([]interface{}, 0, len(insertableImages))
	for i := 0; i < len(insertableImages); i++ {
		m := map[string]interface{}{
			"duration":             aws.Int64Value(insertableImages[i].Duration),
			"fade_in":              aws.Int64Value(insertableImages[i].FadeIn),
			"fade_out":             aws.Int64Value(insertableImages[i].FadeOut),
			"height":               aws.Int64Value(insertableImages[i].Height),
			"image_inserter_input": aws.StringValue(insertableImages[i].ImageInserterInput),
			"image_x":              aws.Int64Value(insertableImages[i].ImageX),
			"image_y":              aws.Int64Value(insertableImages[i].ImageY),
			"layer":                aws.Int64Value(insertableImages[i].Layer),
			"opacity":              aws.Int64Value(insertableImages[i].Opacity),
			"start_time":           aws.StringValue(insertableImages[i].StartTime),
			"width":                aws.Int64Value(insertableImages[i].Width),
		}
		insImgs = append(insImgs, m)
	}
	return insImgs
}
