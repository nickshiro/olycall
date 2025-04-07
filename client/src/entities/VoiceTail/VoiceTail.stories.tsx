import type { Meta, StoryObj } from "@storybook/react";

import { VoiceTail } from "./VoiceTail";

const meta: Meta<typeof VoiceTail> = {
	component: VoiceTail,
	title: "Entities/VoiceTail",
};

export default meta;

type Story = StoryObj<typeof VoiceTail>;

export const Default: Story = {
	args: {
		src: "//m.dedkov.space/meme/ponasenkov",
	},
	decorators: [
		(Story) => (
			<div className="w-46.25">
				<Story />
			</div>
		),
	],
};
