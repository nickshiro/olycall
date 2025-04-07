import { FastAverageColor } from "fast-average-color";
import type { FastAverageColorResult } from "fast-average-color";

const AverageColor = async (
	src: string,
): Promise<FastAverageColorResult | undefined> => {
	const fac = new FastAverageColor();

	try {
		const color = await fac.getColorAsync(src);
		return color;
	} catch (e) {
		console.log(e);
	}
};

export { AverageColor };
