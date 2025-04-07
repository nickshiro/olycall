import type { Meta, StoryObj } from "@storybook/react";

import { VoiceMembers } from "./VoiceMembers";

const meta: Meta<typeof VoiceMembers> = {
	component: VoiceMembers,
	title: "Widgets/VoiceMembers",
};

export default meta;

type Story = StoryObj<typeof VoiceMembers>;

export const Default: Story = {
	args: {},
	decorators: [
		(Story) => (
			<div className="w-89.5">
				<Story />
			</div>
		),
	],
};
