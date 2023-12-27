# Plant Smart
A free service that aims to provide a detailed listing of dangerous plants for your pets.

## Install Locally
Ensure you're using a recent version of `node`. This project was initially built using `19.3.0`. 

```
git clone https://github.com/wilhelm-murdoch/plantsm.art.git
cd plantsm.art
npm install
npm run dev

> plantsm.art@0.0.1 dev
> vite dev


  VITE v4.0.4  ready in 1250 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h to show help
```

Open a browser window and point it to [http://localhost:5173/](http://localhost:5173/). You should now be able to modify anything under `/src` and the local dev build will perform a live update in your browser.

## Updating Datasets
If you want to help out with the dataset, first ensure you have at least Go version `1.19.0` installed locally and run the following commands:
```
cd magefiles/
go mod download
go mod verify
cd ..
mage
```
You should see the following output:
```
Targets:
  images:download
  json:animal
  json:animals
  json:pages
  json:symptoms
  munge:symptoms
  pb:export
  pb:import
```

All persistent data is stored in the primary JSON database in `/static/api/plants.json`. Any changes to this file will require the execution of the following commands:
```
mage json:symptoms static/api/plants.json > static/api/symptoms.json
mage json:animal static/api/plants.json static/api
mage json:animals static/api/plants.json > static/api/animals.json
mage json:slim static/api/plants.json > src/lib/data/slim.json
mage json:pages static/api/plants.json src/routes/plant
mage json:pages static/api/plants.json lib/data/plants
```
A quick explainer of the commands:

* `mage json:symptoms` is used to sync `/static/api/symptoms.json` with any changes `/static/api/plants.json`.
* `mage json:animal` writes individual JSON files for each supported animal.
* `mage json:animals` ( note the plural form ) writes a _single_ JSON document for all supported animals.
* `mage json:pages` writes an individual JSON file for each plant object in `/static/api/plants.json`.
* `mage json:slim` outputs a minified version of `/static/api/plants.json` to be saved as `/src/routes/slim.json`.

### Manage Images

If you've added new plant records that contain new images, be sure to upload them to Cloudflare using the following command:

```
mage images:cloudflare static/api/plants.json
```
This will send new images discovered in `.[].images[].source_url` to Cloudflare which will be found on the path defined in `.[].images[].relative_path`. To access the newly-uploaded image, you simply go to `https://cdn.plantsm.art/cdn-cgi/imagedelivery/qnkf0SBjwoY9e50dCCmY3Q/${relative_path}/(large|medium|thumb)`.