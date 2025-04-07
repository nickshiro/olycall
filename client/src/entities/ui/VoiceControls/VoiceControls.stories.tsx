import type { Meta, StoryObj } from "@storybook/react";

import { VoiceControls } from "./VoiceControls";

const meta: Meta<typeof VoiceControls> = {
	component: VoiceControls,
	title: "Entities/VoiceControl",
};

export default meta;

type Story = StoryObj<typeof VoiceControls>;

export const Default: Story = {
	args: {},
	decorators: [
		(Story) => (
			<div className="w-85.5">
				<Story />
			</div>
		),
	],
};
