import { memo } from "react";
import type { FC } from "react";
import { ToggleButton, Icon } from "@/shared/ui";

const VoiceControls: FC = memo(() => {
	return (
		<nav className="flex items-center gap-x-2 w-full">
			<ToggleButton>
				<div className="w-5 h-5">
					<Icon icon="video" />
				</div>
			</ToggleButton>
			<ToggleButton>
				<div className="w-5 h-5">
					<Icon icon="screencast" />
				</div>
			</ToggleButton>
			<ToggleButton>
				<div className="w-5 h-5">
					<Icon icon="microphone_muted" />
				</div>
			</ToggleButton>
			<ToggleButton>
				<div className="w-5 h-5">
					<Icon icon="headphones" />
				</div>
			</ToggleButton>
			<ToggleButton>
				<div className="w-5 h-5">
					<Icon icon="end_call" />
				</div>
			</ToggleButton>
		</nav>
	);
});

export { VoiceControls };
