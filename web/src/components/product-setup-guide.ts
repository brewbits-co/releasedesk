import {html, LitElement} from 'lit';
import {customElement, state, property} from 'lit/decorators.js';
import {getTailwindStyleSheet} from "../styles";
import {encodeFormData} from "../utils";
import {map} from 'lit/directives/map.js';
import {range} from 'lit/directives/range.js';

@customElement('product-setup-guide')
export class ProductSetupGuide extends LitElement {
    // @ts-ignore
    static styles = [getTailwindStyleSheet()]

    @property({type: String})
    slug = "";

    @state()
    showCustomChannels = false;
    @state()
    totalCustomChannels = 0;

    handleSubmit(event: SubmitEvent) {
        event.preventDefault();

        const formData = new FormData(event.target as HTMLFormElement);

        fetch(`/internal/apps/${this.slug}/setup`, {
            method: 'POST',
            body: encodeFormData(formData),
            credentials: 'include',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        }).then(async (response) => {
            if (!response.ok) {
                alert("error");
            } else {
                window.location.reload();
            }
        });
    }

    handleChannelChange(event: Event) {
        const target = event.target as HTMLInputElement;
        if (target.value == 'CustomChannels') {
            this.showCustomChannels = true;
            this.totalCustomChannels = 1;
        } else {
            this.showCustomChannels = false;
        }
    }

    addNewCustomChannel() {
        this.totalCustomChannels++;
    }

    override render() {
        return html`
            <form @submit="${this.handleSubmit}">
                <ul class="divide-y divide-gray-100 overflow-hidden ring-1 ring-gray-900/5 sm:rounded-xl">
                    <li class="relative flex justify-between gap-x-6 px-4 py-5 sm:px-6">
                        <div class="flex min-w-0 gap-x-4">
                            <div class="relative flex h-5 w-5 flex-shrink-0 items-center justify-center">
                                <div class="h-2 w-2 rounded-full bg-gray-300 group-hover:bg-gray-400"></div>
                            </div>
                            <div class="min-w-0 flex-auto">
                                <p class="text-md font-medium leading-6 text-gray-800 cursor-pointer">
                                    Choose the Version Format
                                </p>
                                <div class="mt-1 text-xs leading-5 text-gray-500">
                                    <fieldset>
                                        <div class="mt-6 grid grid-cols-1 gap-y-6 sm:grid-cols-3 sm:gap-x-4">
                                            ${radioOption('VersionFormat', 'SemVer', 'Semantic Versioning', 'SemVer is a 3-component number in the format of MAJOR.MINOR.PATCH', '1.0.2')}
                                            ${radioOption('VersionFormat', 'CalVer', 'Calendar Versioning', "CalVer is a versioning convention based on your product's release calendar.", '2024.09')}
                                            ${radioOption('VersionFormat', 'Custom', 'Custom Format', "A custom versioning format allows you to fit product-specific needs.", '')}
                                        </div>
                                    </fieldset>
                                </div>
                            </div>
                        </div>
                    </li>
                    <li class="relative flex justify-between gap-x-6 px-4 py-5 sm:px-6">
                        <div class="flex min-w-0 gap-x-4">
                            <div class="relative flex h-5 w-5 flex-shrink-0 items-center justify-center">
                                <div class="h-2 w-2 rounded-full bg-gray-300 group-hover:bg-gray-400"></div>
                            </div>
                            <div class="min-w-0 flex-auto">
                                <p class="text-md font-medium leading-6 text-gray-800 cursor-pointer">
                                    Set Up the Release Channels
                                </p>
                                <div class="mt-1 text-xs leading-5 text-gray-500">
                                    <fieldset>
                                        <div class="mt-6 grid grid-cols-1 gap-y-6 sm:grid-cols-3 sm:gap-x-4">
                                            ${radioOption('Channels', 'ByMaturity', 'By Maturity', 'Channels represent the stability of the release and the intended audience.', 'Canary - Beta - Stable', this.handleChannelChange)}
                                            ${radioOption('Channels', 'ByEnvironment', 'By Environment', 'Channels map releases to environments for effective deployment control.', 'Development - Staging - Production', this.handleChannelChange)}
                                            ${radioOption('Channels', 'CustomChannels', 'Custom Channels', "Create release channels tailored to your product's unique requirements.", '', this.handleChannelChange)}
                                        </div>

                                        ${this.showCustomChannels ? html`
                                            <fieldset name="customChannels" class="mt-4 max-w-md">
                                                <legend class="text-sm font-medium text-gray-900">
                                                    Custom Channels
                                                </legend>
                                                ${map(range(this.totalCustomChannels), () => html`${listItem('CustomChannels', 'Enter the channel name')}`)}
                                                <div class="flex pt-2">
                                                    <button type="button"
                                                            class="text-sm/6 font-semibold text-violet-800 hover:text-violet-700"
                                                            @click="${this.addNewCustomChannel}">
                                                        <span aria-hidden="true">+</span> Add another channel
                                                    </button>
                                                </div>
                                            </fieldset>` : ``}
                                    </fieldset>
                                </div>
                                <button type="submit"
                                        class="rounded-md mt-4 bg-violet-800 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-800">
                                    Save
                                </button>
                            </div>
                        </div>
                    </li>
                </ul>
            </form>
        `;
    }
}

const radioOption = (name: string, value: string, title: string, subtitle: string, footer: string, handler?: any) => html`
    <label class="relative flex cursor-pointer rounded-lg border bg-white p-4 shadow-sm focus:outline-none">
        <input type="radio" name="${name}" value="${value}" class="peer sr-only" @change="${handler}" required>
        <span class="flex flex-1">
            <span class="flex flex-col">
                <span class="block text-sm font-medium text-gray-900">${title}</span>
                <span class="mt-1 flex items-center text-sm text-gray-500">${subtitle}</span>
                <span class="mt-6 text-sm font-medium text-gray-900" translate="no">${footer}</span>
            </span>
        </span>

        <svg class="size-5 text-violet-800 hidden peer-checked:block" viewBox="0 0 20 20" fill="currentColor"
             aria-hidden="true" data-slot="icon">
            <path fill-rule="evenodd"
                  d="M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16Zm3.857-9.809a.75.75 0 0 0-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089l4-5.5Z"
                  clip-rule="evenodd"/>
        </svg>
        <span class="pointer-events-none absolute -inset-px rounded-lg border-2 peer-checked:border-violet-800"
              aria-hidden="true"></span>
    </label>
`

const listItem = (name: string, placeholder: string) => html`
    <div class="divide-y divide-gray-200 border-b border-gray-200">
        <div class="relative flex gap-3 py-2">
            <div class="min-w-0 flex-1 text-sm/6">
                <input type="text"
                       name="${name}"
                       id="${name}"
                       required
                       autocomplete="off"
                       placeholder="${placeholder}"
                       class="block w-full bg-white px-3 py-1.5 text-base text-gray-900 placeholder:text-gray-400 outline-0 border-0 focus:ring-0 sm:text-sm/6">
            </div>
        </div>
    </div>
`;

declare global {
    interface HTMLElementTagNameMap {
        'product-setup-guide': ProductSetupGuide;
    }
}
