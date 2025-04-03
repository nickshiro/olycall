import type { FC } from "react";
import type { Icon } from "../Icon.type";
import {
	Headphones,
	HeadphonesMuted,
	Leave,
	MicrophoneMuted,
	Screencast,
	Settings,
	Video,
} from "../icons";
import { ArrowUp } from "../icons/ArrowUp";

const icons: Record<Icon, FC> = {
	leave: Leave,
	headphones: Headphones,
	headphones_muted: HeadphonesMuted,
	microphone_muted: MicrophoneMuted,
	video: Video,
	screencast: Screencast,
	settings: Settings,
	arrow_up: ArrowUp,
};

export { icons };
