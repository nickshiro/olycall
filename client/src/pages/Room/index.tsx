import { Panel, VoiceMembers } from "@/widgets";
import { useState, type FC } from "react";
import { TheatreMode } from "./ui/TheatreMode";
import { TailsMode } from "./ui/TailsMode";

const Room: FC = () => {
	const [isTailsMode, setIsTailsMode] = useState<boolean>(false);

	return (
		<div className="flex w-full p-4 box-border h-screen flex-1">
			<div className="flex box-border gap-x-2 m-auto w-full">
				<div className="flex flex-col justify-between max-w-89.5 lg:w-89.5">
					<VoiceMembers />
					<Panel />
				</div>
				<div className="bg-bg-secondary w-full max-h-full p-4 box-border rounded-2xl aspect-video">
					{isTailsMode ? (
						<TailsMode setTailsMode={setIsTailsMode} />
					) : (
						<TheatreMode setTailsMode={setIsTailsMode} />
					)}
				</div>
			</div>
		</div>
	);
};

export default Room;
