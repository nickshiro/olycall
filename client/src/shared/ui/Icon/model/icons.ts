import type { FC } from "react";

// Icons
import Video from "./icons/Video.svg?react";
import MicrophoneMuted from "./icons/Microphone_muted.svg?react";
import Screencast from "./icons/Screencast.svg?react";
import Apps from "./icons/Apps.svg?react";
import EndCall from "./icons/End_call.svg?react";
import Settings from "./icons/Settings.svg?react";
import Expand from "./icons/Expand.svg?react";
import Headphones from "./icons/Headphones.svg?react";

type IIcons =
	| "headphones"
	| "microphone_muted"
	| "screencast"
	| "video"
	| "apps"
	| "end_call"
	| "settings"
	| "expand";

const icons: Record<IIcons, FC> = {
	video: Video,
	microphone_muted: MicrophoneMuted,
	screencast: Screencast,
	apps: Apps,
	end_call: EndCall,
	settings: Settings,
	expand: Expand,
	headphones: Headphones,
};

export type { IIcons };
export { icons };
