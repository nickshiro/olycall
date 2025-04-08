import { memo } from "react";
import type { FC } from "react";

import { VoiceMember } from "@/entities/ui";

const VoiceMembers: FC = memo(() => {
	return (
		<div className="rounded-2xl p-2 box-border flex flex-col gap-y-1 bg-bg-secondary">
			<VoiceMember src="//m.dedkov.space/meme/ponasenkov" />
			<VoiceMember src="//m.dedkov.space/meme/ponasenkov" />
			<VoiceMember src="//m.dedkov.space/meme/ponasenkov" />
			<VoiceMember src="//m.dedkov.space/meme/ponasenkov" />
		</div>
	);
});

export { VoiceMembers };
