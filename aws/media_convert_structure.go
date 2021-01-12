package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

func expandMediaConvertReservationPlanSettings(config map[string]interface{}) *mediaconvert.ReservationPlanSettings {
	reservationPlanSettings := &mediaconvert.ReservationPlanSettings{}

	if v, ok := config["commitment"]; ok {
		reservationPlanSettings.Commitment = aws.String(v.(string))
	}

	if v, ok := config["renewal_type"]; ok {
		reservationPlanSettings.RenewalType = aws.String(v.(string))
	}

	if v, ok := config["reserved_slots"]; ok {
		reservationPlanSettings.ReservedSlots = aws.Int64(int64(v.(int)))
	}

	return reservationPlanSettings
}

func flattenMediaConvertReservationPlan(reservationPlan *mediaconvert.ReservationPlan) []interface{} {
	if reservationPlan == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"commitment":     aws.StringValue(reservationPlan.Commitment),
		"renewal_type":   aws.StringValue(reservationPlan.RenewalType),
		"reserved_slots": aws.Int64Value(reservationPlan.ReservedSlots),
	}

	return []interface{}{m}
}

func expandMediaConvertNexGuardFileMarkerSettings(list []interface{}) *mediaconvert.NexGuardFileMarkerSettings {
	result := &mediaconvert.NexGuardFileMarkerSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertDolbyVisionLevel6Metadata(list []interface{}) *mediaconvert.DolbyVisionLevel6Metadata {
	result := &mediaconvert.DolbyVisionLevel6Metadata{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertNoiseReducerTemporalFilterSettings(list []interface{}) *mediaconvert.NoiseReducerTemporalFilterSettings {
	result := &mediaconvert.NoiseReducerTemporalFilterSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertNoiseReducerFilterSettings(list []interface{}) *mediaconvert.NoiseReducerFilterSettings {
	result := &mediaconvert.NoiseReducerFilterSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerSpatialFilterSettings(list []interface{}) *mediaconvert.NoiseReducerSpatialFilterSettings {
	result := &mediaconvert.NoiseReducerSpatialFilterSettings{}
	if list == nil || len(list) == 0 {
		return nil
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
	if v, ok := tfMap["partner_watermarking"]; ok {
		result.PartnerWatermarking = expandMediaConvertPartnerWatermarking(v.([]interface{}))
	}
	if v, ok := tfMap["timecode_burnin"]; ok {
		result.TimecodeBurnin = expandMediaConvertTimecodeBurnin(v.([]interface{}))
	}
	return result
}

func expandMediaConvertNoiseReducer(list []interface{}) *mediaconvert.NoiseReducer {
	result := &mediaconvert.NoiseReducer{}
	if list == nil || len(list) == 0 {
		return nil
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
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["nexguard_file_marker_settings"]; ok {
		result.NexguardFileMarkerSettings = expandMediaConvertNexGuardFileMarkerSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertTimecodeBurnin(list []interface{}) *mediaconvert.TimecodeBurnin {
	result := &mediaconvert.TimecodeBurnin{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertDolbyVision(list []interface{}) *mediaconvert.DolbyVision {
	result := &mediaconvert.DolbyVision{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertDeinterlacer(list []interface{}) *mediaconvert.Deinterlacer {
	result := &mediaconvert.Deinterlacer{}
	if list == nil || len(list) == 0 {
		return nil
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
		return nil
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

func expandMediaConvertVideoCodecSettings(list []interface{}) *mediaconvert.VideoCodecSettings {
	result := &mediaconvert.VideoCodecSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertAv1QvbrSettings(list []interface{}) *mediaconvert.Av1QvbrSettings {
	result := &mediaconvert.Av1QvbrSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertH264QvbrSettings(list []interface{}) *mediaconvert.H264QvbrSettings {
	result := &mediaconvert.H264QvbrSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func expandMediaConvertH265QvbrSettings(list []interface{}) *mediaconvert.H265QvbrSettings {
	result := &mediaconvert.H265QvbrSettings{}
	if list == nil || len(list) == 0 {
		return nil
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

func flattenMediaConvertAudioChannelTaggingSettings(cfg *mediaconvert.AudioChannelTaggingSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"channel_tag": aws.StringValue(cfg.ChannelTag),
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

func flattenMediaConvertVideoPreprocessor(cfg *mediaconvert.VideoPreprocessor) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"color_corrector":      flattenMediaConvertColorCorrector(cfg.ColorCorrector),
		"deinterlacer":         flattenMediaConvertDeinterlacer(cfg.Deinterlacer),
		"dolby_vision":         flattenMediaConvertDolbyVision(cfg.DolbyVision),
		"image_inserter":       flattenMediaConvertImageInserter(cfg.ImageInserter),
		"noise_reducer":        flattenMediaConvertNoiseReducer(cfg.NoiseReducer),
		"partner_watermarking": flattenMediaConvertPartnerWatermarking(cfg.PartnerWatermarking),
		"timecode_burnin":      flattenMediaConvertTimecodeBurnin(cfg.TimecodeBurnin),
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
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"nexguard_file_marker_settings": flattenMediaConvertNexGuardFileMarkerSettings(cfg.NexguardFileMarkerSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertNexGuardFileMarkerSettings(cfg *mediaconvert.NexGuardFileMarkerSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"license":  aws.StringValue(cfg.License),
		"payload":  aws.Int64Value(cfg.Payload),
		"preset":   aws.StringValue(cfg.Preset),
		"strength": aws.StringValue(cfg.Strength),
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
