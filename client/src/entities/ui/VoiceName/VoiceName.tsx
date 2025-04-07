import { memo } from "react";
import type { FC } from "react";

export interface VoiceNameProps {
	children: string;
}

const VoiceName: FC<VoiceNameProps> = memo(({ children }) => {
	return (
		<h1 className="text-text-primary text-base font-medium select-none">
			{children}
		</h1>
	);
});

export { VoiceName };
