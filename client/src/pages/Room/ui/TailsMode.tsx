import { memo } from "react";
import type { FC } from "react";
import { VoiceTail } from "@/entities/ui/VoiceTail/VoiceTail";

interface TailsModeProps {
	setTailsMode: (x: boolean) => void;
}

const TailsMode: FC<TailsModeProps> = memo(() => {
	return (
		<div className="w-full h-full flex flex-col justify-center">
			<div className="grid grid-cols-3 gap-2">
				<VoiceTail
					className="aspect-video"
					src="//m.dedkov.space/meme/ponasenkov"
				/>
				<VoiceTail
					className="aspect-video"
					src="//m.dedkov.space/meme/misha_stoit"
				/>
				<VoiceTail
					className="aspect-video"
					src="//m.dedkov.space/meme/ponasenkov"
				/>
				<VoiceTail
					className="aspect-video"
					src="//m.dedkov.space/meme/misha_stoit"
				/>
				<VoiceTail
					className="aspect-video"
					src="//m.dedkov.space/meme/ponasenkov"
				/>
			</div>
		</div>
	);
});

export { TailsMode };
