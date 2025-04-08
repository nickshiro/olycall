import { memo, useState } from "react";
import type { FC } from "react";

import { Icon, Button } from "@/shared/ui";
import { VoiceAvatar, VoiceControls, VoiceName } from "@/entities/ui";

import { JoinButton } from "./ui/JoinButton";

const Panel: FC = memo(() => {
	const [isJoined, _] = useState<boolean>(true);

	return (
		<div className="w-full rounded-2xl bg-bg-secondary p-2 flex flex-col gap-y-2">
			<section className="py-1.5 px-3 box-border flex items-center justify-between w-full">
				<div className="flex items-center gap-x-4">
					<VoiceAvatar src="//m.dedkov.space/meme/ponasenkov" isActive />
					<VoiceName>Pavel Durov</VoiceName>
				</div>
				<div className="h-8 w-8 ml-2">
					<Button>
						<div className="w-5 h-5">
							<Icon icon="settings" />
						</div>
					</Button>
				</div>
			</section>
			<hr className="text-[var(--dark-850)]" />
			{isJoined ? <VoiceControls /> : <JoinButton />}
		</div>
	);
});

export { Panel };
