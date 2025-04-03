import { Avatar, Icon, Name } from "@/shared/ui";
import type { FC } from "react";
import { memo } from "react";
import { IconWrapper } from "./ui/IconWrapper";

export interface MemberProps {
	avatar: string;
	name: string;
	isMuted?: boolean;
	isDeaf?: boolean;
	isScreencast?: boolean;
	isVideo?: boolean;
	isActive?: boolean;
}

const MemberComponent: FC<MemberProps> = ({
	avatar,
	name,
	isMuted = false,
	isDeaf = false,
	isScreencast = false,
	isVideo = false,
	isActive = false,
}) => {
	return (
		<div className="flex items-center justify-between hover:bg-bg-quaternary cursor-pointer px-2 py-1 rounded-lg transition-colors duration-200">
			<section className="flex items-center gap-x-2">
				<Avatar src={avatar} alt={name} isActive={isActive} />
				<Name>{name}</Name>
			</section>
			<ul className="flex items-center gap-x-0.5">
				{isMuted && (
					<IconWrapper>
						<Icon icon="microphone_muted" />
					</IconWrapper>
				)}
				{isDeaf && (
					<IconWrapper>
						<Icon icon="headphones_muted" />
					</IconWrapper>
				)}
				{isScreencast && (
					<IconWrapper>
						<Icon icon="screencast" />
					</IconWrapper>
				)}
				{isVideo && (
					<IconWrapper>
						<Icon icon="video" />
					</IconWrapper>
				)}
			</ul>
		</div>
	);
};

export const Member = memo(MemberComponent);
