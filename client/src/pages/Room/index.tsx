import { VoiceTail } from "@/entities/ui/VoiceTail/VoiceTail";
import { Button, Icon } from "@/shared/ui";
import { Panel, VoiceMembers } from "@/widgets";
import type { FC } from "react";

const Room: FC = () => {
	return (
		<div className="flex h-screen items-center w-full box-border">
			<div className="p-4 flex gap-x-2 box-border w-full">
				<div className="w-89.5 flex flex-col justify-between">
					<VoiceMembers />
					<Panel />
				</div>
				<div className="bg-bg-secondary p-4 rounded-2xl aspect-video w-full box-border flex flex-col justify-stretch gap-y-4">
					<div className="">
						<div className="flex gap-x-4">
							<div className="w-8 h-8">
								<Button className="bg-bg-tertiary hover:bg-bg-quaternary">
									<div className="w-5 h-5">
										<Icon icon="apps" />
									</div>
								</Button>
							</div>
							<div className="flex flex-1 aspect-video">
								<div className="aspect-video w-full box-border bg-accent-primary rounded-lg" />
							</div>
							<div className="w-8 h-8">
								<Button className="bg-bg-tertiary hover:bg-bg-quaternary">
									<div className="w-5 h-5">
										<Icon icon="expand" />
									</div>
								</Button>
							</div>
						</div>
					</div>
					<div className="flex gap-x-2 h-26 justify-center box-border">
						<VoiceTail
							className="hover:cursor-pointer"
							src="//m.dedkov.space/meme/ponasenkov"
						/>
						<VoiceTail
							className="hover:cursor-pointer"
							src="//m.dedkov.space/meme/ponasenkov"
						/>
						<VoiceTail
							className="hover:cursor-pointer"
							src="//m.dedkov.space/meme/ponasenkov"
						/>
						<VoiceTail
							className="hover:cursor-pointer"
							src="//m.dedkov.space/meme/ponasenkov"
						/>
					</div>
				</div>
			</div>
		</div>
	);
};

export default Room;
