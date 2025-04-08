import { memo } from "react";
import type { FC } from "react";
import { VoiceTail } from "@/entities/ui/VoiceTail/VoiceTail";
import { Button, Icon } from "@/shared/ui";

interface TheatreModeProps {
	setTailsMode: (x: boolean) => void;
}

const TheatreMode: FC<TheatreModeProps> = memo(({ setTailsMode }) => {
	return (
		<div className="flex flex-col gap-y-4 box-border h-full">
			<div className="flex flex-1 gap-x-4">
				<div className="w-8">
					<div className="h-8 w-8">
						<Button
							className="bg-bg-tertiary hover:bg-bg-quaternary"
							onClick={() => setTailsMode(true)}
						>
							<div className="h-5 w-5">
								<Icon icon="apps" />
							</div>
						</Button>
					</div>
				</div>
				<div className="flex-1 flex items-stretch justify-stretch box-border">
					<div className="aspect-video bg-accent-primary mx-auto rounded-lg">
						gg
					</div>
				</div>
				<div className="w-8">
					<div className="h-8 w-8">
						<Button className="bg-bg-tertiary hover:bg-bg-quaternary">
							<div className="h-5 w-5">
								<Icon icon="expand" />
							</div>
						</Button>
					</div>
				</div>
			</div>
			<div className="h-26 gap-x-2 flex w-full justify-center overflow-x-auto">
				<VoiceTail
					className="max-w-46.25"
					src="//m.dedkov.space/meme/ponasenkov"
				/>
				<VoiceTail className="max-w-46.25" src="//m.dedkov.space/meme/nagiev" />
				<VoiceTail
					className="max-w-46.25"
					src="//m.dedkov.space/meme/ponasenkov"
				/>
				<VoiceTail
					className="max-w-46.25"
					src="//m.dedkov.space/meme/misha_stoit"
				/>
				<VoiceTail
					className="max-w-46.25"
					src="//m.dedkov.space/ponasenkov/elegant"
				/>
				<VoiceTail
					className="max-w-46.25"
					src="//m.dedkov.space/meme/misha_stoit"
				/>
			</div>
		</div>
	);
});

export { TheatreMode };
