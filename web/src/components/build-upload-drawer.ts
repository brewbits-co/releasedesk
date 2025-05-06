import {html, LitElement} from 'lit';
import {customElement, property, state} from 'lit/decorators.js';
import {createRef, Ref, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

@customElement('build-upload-drawer')
export class BuildUploadDrawer extends LitElement {
    override createRenderRoot() {
        return this;
    }

    extensionsByPlatform = {
        "Windows": [".appx", ".appxbundle", ".appxupload", ".msix", ".msixbundle", ".msixupload", ".exe", ".zip", ".msi"],
        "macOS": [".zip", ".platform.zip", ".dmg", ".pkg"],
        "Linux": [".AppImage", ".deb", ".rpm",  ".tgz", ".gz", ".snap", ".flatpak"],
        "Android": [".apk", ".aab"],
        "iOS": [".ipa"],
    };

    archsByPlatform = {
        "Windows": ["x86", "x64"],
        "macOS": ["x64", "ARM64"],
        "Linux": ["x86", "x64", "ARM", "ARM64"],
        "Android": [""],
        "iOS": [""],
    };

    @property({type: String})
    slug = "";

    @property({type: String})
    platform = "Windows";

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

        fetch(`/internal/apps/${this.slug}/platforms/${this.platform}/builds`, {
            method: 'POST',
            body: formData,
            credentials: 'include'
        }).then(async response => {
            if (!response.ok) {
                this.submitFailed = true;
                this.errorData = await response.json();
                // wait form render to complete
                await this.updateComplete;
                // scroll the page to the error alert
                this.errorAlertRef.value?.scrollIntoView();
            } else {
                window.location.href = `/dashboard/${this.slug}/platforms/${this.platform}/builds/${formData.get("Number")}`;
            }
        });
    }

    handleClose() {
        this.showDrawer = false;
        this.submitFailed = false;
    }

    override render() {
        return html`
            <button type="button"
                    class="inline-flex items-center rounded-md bg-violet-900 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-900"
                    @click="${() => this.showDrawer = true}">
                <i class="fa-solid fa-upload"></i>&nbsp;&nbsp;Upload
            </button>

            <headless-drawer ?open="${this.showDrawer}"
                             @closed="${this.handleClose}"
                             headerTitle="Upload Build">
                <form slot="body" id="build-upload-form" class="px-4 sm:px-6 pb-4"
                      @reset="${this.handleClose}"
                      @submit="${this.handleSubmit}">
                    <div class="space-y-6 pb-5 pt-6">
                        <div>
                            <label for="project-name" class="block text-sm/6 font-medium text-gray-900">Number</label>
                            <div class="mt-2">
                                <input type="text"
                                       name="Number"
                                       id="Number"
                                       class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                       maxlength="255"
                                       autocomplete="off"
                                       required>
                            </div>
                        </div>
                        <div>
                            <label for="project-name" class="block text-sm/6 font-medium text-gray-900">Version</label>
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
                            ${
                                    // @ts-ignore 
                                    map(this.archsByPlatform[this.platform], (arch) => html`
                                        <div class="mt-2">
                                            <label for="File${arch ? '_' + arch : ''}"
                                                   class="block text-sm/6 font-medium text-gray-900">Artifact
                                                ${arch}</label>
                                            <input type="file"
                                                   name="File${arch ? '_' + arch : ''}"
                                                   id="File${arch ? '_' + arch : ''}"
                                                   accept="${
                                                           // @ts-ignore 
                                                           this.extensionsByPlatform[this.platform].join(',')
                                                   }"
                                                   class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                                   maxlength="255"
                                                   autocomplete="off">
                                            <p class="text-xs/5 text-gray-600">Upload ${
                                                    // @ts-ignore 
                                                    this.extensionsByPlatform[this.platform].join(', ')
                                            } file(s)</p>
                                        </div>`
                                    )}
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
                  form="build-upload-form"
                  class="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
            Cancel
          </button>
          <button type="submit"
                  form="build-upload-form"
                  class="ml-4 inline-flex justify-center rounded-md bg-violet-800 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-800">
            Upload
          </button>
        </span>
            </headless-drawer>
        `;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        'build-upload-drawer': BuildUploadDrawer;
    }
}
