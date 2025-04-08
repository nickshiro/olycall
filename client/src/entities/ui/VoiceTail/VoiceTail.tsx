import { AverageColor } from "@/shared/lib";
import { Avatar } from "@/shared/ui";
import { memo, useEffect, useState } from "react";
import type { FC } from "react";
import cn from "clsx";

export interface VoiceTailProps {
	src?: string;
	className?: string;
}

const VoiceTail: FC<VoiceTailProps> = memo(({ src = "", className }) => {
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

	const classes = cn(
		className,
		"rounded-lg h-full box-border flex items-center justify-center p-4 box-border h-full w-full",
	);

	return (
		<div className={classes} style={{ backgroundColor: color ? color : "" }}>
			<Avatar src={src} size="small" />
		</div>
	);
});

export { VoiceTail };
