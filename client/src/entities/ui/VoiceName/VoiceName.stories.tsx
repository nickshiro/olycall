import type { Meta, StoryObj } from "@storybook/react";

import { VoiceName } from "./VoiceName";

const meta: Meta<typeof VoiceName> = {
	component: VoiceName,
	title: "Entities/VoiceName",
};

export default meta;

type Story = StoryObj<typeof VoiceName>;

export const Default: Story = {
	args: {
		children: "Pavel Durov",
	},
};
