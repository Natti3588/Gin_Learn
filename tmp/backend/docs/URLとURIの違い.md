# URLとURIの違い

**結論**は「**URIが広い概念で、URLはその一種**」

## URI（Uniform Resource Identifier）
リソース（画像、JSON）を識別するための文字列の総称。 
識別子なので、何かを一位に指し示せれば、それは**URI**

## URL（Uniform Resource Locator）
**URI**のうち、「どこにあるのか」をリソースで指すものです。
`https://example.com/page`のように、プロトコル + 場所が含まれて、
実際にそこへアクセスできます。

## 関係のイメージ
```
URI（識別子すべて）
├── URL（場所で指す）   例: https://example.com/index.html
└── URN（名前で指す）   例: urn:isbn:9784774142230
```
