# regsan-check
Simple go API to check for the existence of NSO codes in a google sheets database.

## NSO and sanitary registry in Ecuador
Sanitary Obligatory Notification (NSO) is the main government requirement to sell cosmetic, health and food products in Ecuador. It is also known Sanitary Registry and it is issued by the sanitary authority ([ARCSA](https://www.controlsanitario.gob.ec)).

## Counterfeit products 
NSO code must be shown in the product label, therfore, false and counterfeit productos include false NSO codes or no codes at all. To verify the validy of a NSO code, buyers and consumers can check the official [ARCSA data bases](https://www.controlsanitario.gob.ec/base-de-datos/). Since the government agency's webpage may be hard to browse, a simple api endpoint may help to make this information easily accesible to the public.

## Usage
Open this [example](https://regsan-check.herokuapp.com/check-nso/NSOC07248-11P/) in your browser and change the last part in the url to the NSO code in the product you would like to check.

```https://regsan-check.herokuapp.com/check-nso/<product NSO code>/```

## The api
Written in Go (Golang) as a learning (toy) project and deployed in Heroku. It uses a Google Service Account (gserviceaccount) to access data in a Google Spreadsheet and serve the data in json format.
**Do not use this endpoint or data for legal or official purposes. Since this is a learning project, data is not guaranteed to be veracius or updated. If you really doubt the truthfulness of a product consult the official government agency [ARCSA](https://www.controlsanitario.gob.ec)**

## Data
Currently, only cosmetic products data is included in a mirror database used for backend of this api. The api response, in case of existing NSO codes includes these variables:

* NSO: product NSO code or sanitary registry
* NombreProducto: Product name
* MarcaProducto: Brand name
* Titular: NSO holder (person or company)
* FechaEmision: Issue date
* FechaVigencia: Validity date
