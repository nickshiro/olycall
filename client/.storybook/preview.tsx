import React from "react";
import type { Preview } from "@storybook/react";

// Styles
import "@/app/styles/global.css";
import "@/shared/fonts/fonts.css";

const preview: Preview = {
	parameters: {
		controls: {
			matchers: {
				color: /(background|color)$/i,
				date: /Date$/i,
			},
		},
		layout: "fullscreen",
	},
	decorators: [
		(Story) => (
			<div className="flex items-center justify-center min-h-screen box-border">
				<Story />
			</div>
		),
	],
};

export default preview;
