import { Avatar } from "@/shared/ui";
import { memo } from "react";
import type { FC } from "react";

export interface VoiceAvatarProps {
	src?: string;
	isActive?: boolean;
}

const VoiceAvatar: FC<VoiceAvatarProps> = memo(({ src, isActive = false }) => {
	return (
		<div className="h-10 w-10 relative transition-colors">
			{isActive && (
				<div className="absolute w-full h-full rounded-full outline outline-accent-primary border border-bg-secondary" />
			)}
			<Avatar src={src ? src : ""} size="small" />
		</div>
	);
});

export { VoiceAvatar };
