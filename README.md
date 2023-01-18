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

All persistent data is stored in the primary JSON database in `/static/plants.json`. Any changes to this file will require the execution of the following commands:
```
mage json:symptoms static/plants.json > static/symptoms.json
mage json:animal static/plants.json static
mage json:animals static/plants.json > static/animals.json
mage json:slim static/plants.json > src/routes/slim.json
mage json:pages static/plants.json static/plant
mage json:pages static/plants.json src/routes/plant
```
A quick explainer of the commands:

* `mage json:symptoms` is used to sync `/static/symptoms.json` with any changes `/static/plants.json`.
* `mage json:animal` writes individual JSON files for each supported animal.
* `mage json:animals` ( note the plural form ) writes a _single_ JSON document for all supported animals.
* `mage json:pages` writes an individual JSON file for each plant object in `/static/plants.json`.
* `mage json:slim` outputs a minified version of `/static/plants.json` to be saved as `/src/routes/slim.json`.

### Downloading Images

If you've added new plant records that require image downloads be sure to run the following commands:

```
mage images:download static/plants.json static/images
for image in ./**/large.jpg; do convert "$image" -resize 50% "$(dirname $image)/medium.jpg"; done
for image in ./**/large.jpg; do convert "$image" -resize 192x192 "$(dirname $image)/thumbnail.jpg"; done
```
This will download new images from `.[].images[].source_url` to the path defined in `.[].images[].relative_path`. The subsequent loops use ImageMagick's `convert` tool to create medium-sized and thumbnail versions of the original image.

### Using PocketBase
Personally, I use [PocketBase](https://pocketbase.io/) to update `/static/plants.json`. Ensure you have [Docker](https://www.docker.com/), or [podman](https://podman.io/) if you prefer, installed and running locally. Use `docker-compose` to start the service:
```
docker-compose up
[+] Running 1/0
 ⠿ Container pocketbase  Created  0.0s
Attaching to pocketbase
pocketbase  | > Server started at: http://0.0.0.0:8090
pocketbase  |   - REST API: http://0.0.0.0:8090/api/
pocketbase  |   - Admin UI: http://0.0.0.0:8090/_/
```
Open [http://0.0.0.0:8090/_/](http://0.0.0.0:8090/_/) in your browser. If this is your first time running the service, you will be asked to create a new admin account. Create something memorable and sign in. Go to "Settings > Import Collections" and paste the body of `/data/pb_schema.json` into the textarea. Click "Review" and "Confirm and import" on the following screen.

Next, you will prime this new database with the data from `/static/plants.json`. In your terminal, run the following command:
```
mage pb:import /static/plants.json
```
This will populate the database. Once complete, you can go back to the admin panel and modify the records using a nice interface. When finished, be sure to run the following command to rebuild the `/static/plants.json` file:
```
mage pb:export /static/plants.json
```
This will sync your all your changes. Be sure to run the previously-mentioned `mage` commands outlined in the first section under [Updating Datasets](#updating-datasets).

All changes to any `*.json` database will be immediately reflected in your local dev environment.

