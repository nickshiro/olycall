import { AverageColor } from "@/shared/lib";
import { Avatar } from "@/shared/ui";
import { memo, useEffect, useState } from "react";
import type { FC } from "react";

export interface VoiceTailProps {
	src?: string;
}

const VoiceTail: FC<VoiceTailProps> = memo(({ src = "" }) => {
	const [color, setColor] = useState<string | null>(null);

	useEffect(() => {
		const fetchColor = async () => {
			const colorData = await AverageColor(src);
			if (colorData) {
				setColor(colorData.hex);
			}
		};

		if (src) {
			fetchColor();
		}
	}, [src]);

	return (
		<div
			className="rounded-lg box-border aspect-[16/9] w-full flex items-center justify-center"
			style={{ backgroundColor: color ? color : "" }}
		>
			<Avatar src={src} size="small" />
		</div>
	);
});

export { VoiceTail };
