interface ViteTypeOptions {
  strictImportMetaEnv: unknown;
}

interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  readonly VITE_CANONICAL_URL: string;
  readonly VITE_TITLE_IMAGE_URL: string;
  readonly VITE_PREVIEW_IMAGE_URL: string;
  readonly VITE_CONFIRMATION_IMAGE_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
