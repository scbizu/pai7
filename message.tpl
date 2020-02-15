Build finished for *{{tpl.app}}*

{{#success build.status}}
  🎉 build {{build.number}} succeeded.
{{else}}
  ❌ build {{build.number}} failed.
{{/success}}
