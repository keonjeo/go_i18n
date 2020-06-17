# Golang I18n

Golang internationalization and localization (i18n) solution.

Currently maintained by @keonjeo.

## Usage

If you want to use this library with GOMODULE, you can simply add `i18n` to your `pkgset`:

```bash
go get -u github.com/keonjeo/i18n
```

Then configure I18n with some translations, and a default locale:

```Golang
I18n.DefaultLocale = "en" # (note that `en` is already the default!)
```

A simple translation file in your project might live at `config/locales/en.yml` and look like:

```yml
en:
  test: "This is a test"
```

You can then access this translation by doing:

```golang
I18n.T("test")
```

You can switch locales in your project by setting `I18n.Locale` to a different value:

```Golang
I18n.Locale = "de"
I18n.T("test") # => "Dies ist ein Test"
```

## Contributors

* @keonjeo
* [and many more](https://github.com/keonjeo/i18n/graphs/contributors)

## License

MIT License. See the included MIT-LICENSE file.
