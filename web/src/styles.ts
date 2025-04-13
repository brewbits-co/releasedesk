let globalSheets: CSSStyleSheet[] | null = null;

export function getTailwindStyleSheet() {
  if (globalSheets === null) {
    globalSheets = Array.from(document.styleSheets)
      .filter(x => x.href && x.href.includes("tailwind.css"))
      .map(x => {
        const sheet = new CSSStyleSheet();
        const css = Array.from(x.cssRules).map(rule => rule.cssText).join(' ');
        sheet.replaceSync(css);
        return sheet;
      });
  }

  return globalSheets;
}
