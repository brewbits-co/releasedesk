import {html, LitElement} from 'lit';
import {customElement, property, state} from 'lit/decorators.js';
import {encodeFormData} from "../utils";
import {Ref, createRef, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

@customElement('release-creation-drawer')
export class ReleaseCreationDrawer extends LitElement {
    override createRenderRoot() {
        return this;
    }

    @property({type: String})
    slug = "";

    @property({type: Array})
    channels: Channel[] = [];

    @property({type: Boolean})
    showDrawer = false;

    @state()
    submitFailed = false;

    errorAlertRef: Ref<HTMLDivElement> = createRef();

    @state()
    errorData: ErrorResponse = {message: '', helpTexts: []};

    handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        this.submitFailed = false;

        const formData = new FormData(event.target as HTMLFormElement);

        fetch(`/internal/apps/${this.slug}/releases`, {
            method: 'POST',
            body: encodeFormData(formData),
            credentials: 'include',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        }).then(async response => {
            if (!response.ok) {
                this.submitFailed = true;
                this.errorData = await response.json();
                // wait form render to complete
                await this.updateComplete;
                // scroll the page to the error alert
                this.errorAlertRef.value?.scrollIntoView();
            } else {
                let data = await response.json();
                window.location.href = `/dashboard/${this.slug}/releases/${data.Version}`;
            }
        });
    }

    handleClose() {
        this.showDrawer = false;
        this.submitFailed = false;
    }

    override render() {
        return html`
            <div class="text-center">
                <button type="button"
                        class="inline-flex items-center rounded-md bg-violet-900 gap-x-2 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-900"
                        @click="${() => this.showDrawer = true}">
                    <i class="fa-solid fa-rocket"></i> Create
                </button>
            </div>

            <headless-drawer ?open="${this.showDrawer}"
                             @closed="${this.handleClose}"
                             headerTitle="New Application">
                <form slot="body" id="product-creation-form" class="px-4 sm:px-6 pb-4"
                      @reset="${this.handleClose}"
                      @submit="${this.handleSubmit}">
                    <div class="space-y-6 pb-5 pt-6">
                        <div>
                            <label for="Version" class="block text-sm/6 font-medium text-gray-900">Version</label>
                            <div class="mt-2">
                                <input type="text"
                                       name="Version"
                                       id="Version"
                                       class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                       maxlength="255"
                                       autocomplete="off"
                                       required>
                            </div>
                        </div>
                        <div>
                            <label for="TargetChannel" class="block text-sm/6 font-medium text-gray-900">Channel</label>
                            <div class="mt-2 grid grid-cols-1">
                                <select id="TargetChannel" name="TargetChannel"
                                        class="col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pl-3 pr-8 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-violet-600 sm:text-sm/6">
                                    ${map(this.channels, (channel) => html`
                                        <option translate="no" value="${channel.ID}">${channel.Name}</option>
                                    `)}
                                </select>
                            </div>
                        </div>
                    </div>

                    ${this.submitFailed ? html`
                        <error-alert ${ref(this.errorAlertRef)}>
                            <span slot="message">${this.errorData.message}</span>
                            <ul role="list" class="list-disc space-y-1 pl-5 mt-2" slot="description">
                                ${map(this.errorData.helpTexts, (text) => html`
                                    <li>${text}</li>`
                                )}
                            </ul>
                        </error-alert>` : ``}
                </form>

                <span slot="footer">
          <button type="reset"
                  form="product-creation-form"
                  class="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
            Cancel
          </button>
          <button type="submit"
                  form="product-creation-form"
                  class="ml-4 inline-flex justify-center rounded-md bg-violet-800 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-800">
            Save
          </button>
        </span>
            </headless-drawer>
        `;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        'release-creation-drawer': ReleaseCreationDrawer;
    }
}
