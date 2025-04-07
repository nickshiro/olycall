import { memo } from "react";
import type { FC } from "react";
import { VoiceName } from "../VoiceName";
import { Icon } from "@/shared/ui";
import { VoiceAvatar } from "../VoiceAvatar";

export interface VoiceMemberProps {
	isMuted?: boolean;
	isDeaf?: boolean;
	isScreenshare?: boolean;
	isVideo?: boolean;
	src?: string;
}

const VoiceMember: FC<VoiceMemberProps> = memo(
	({ isMuted, isDeaf, isScreenshare, isVideo, src }) => {
		return (
			<div className="w-full flex items-center justify-between hover:bg-bg-tertiary hover:cursor-pointer rounded-lg py-1.5 px-3">
				<div className="flex items-center gap-x-4">
					<VoiceAvatar src={src} isActive />
					<VoiceName>Pavel Durov</VoiceName>
				</div>
				<ul className="flex items-center text-accent-primary">
					{isMuted && (
						<li className="h-5 w-5">
							<Icon icon="microphone_muted" />
						</li>
					)}
					{isDeaf && (
						<li className="h-5 w-5">
							<Icon icon="headphones" />
						</li>
					)}
					{isVideo && (
						<li className="h-5 w-5">
							<Icon icon="video" />
						</li>
					)}
					{isScreenshare && (
						<li className="h-5 w-5">
							<Icon icon="screencast" />
						</li>
					)}
				</ul>
			</div>
		);
	},
);

export { VoiceMember };
