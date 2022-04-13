# File

File Requests start with a `file` block which specifies the start of the requests for the File template.

The fields supported by File requests are specified with details below.

## extensions

Extensions contains a list of extensions to perform matching on.

```yaml
# "all" means match all extensions except denylist
extensions:
  - "all"
```

A custom list of extensions to be matched on can also be provided.

```yaml
# Custom extension list
extensions:
  - .txt
  - .go
  - .json
```

## denylist

Denylist is the list of extensions to deny matching upon. By default, nuclei file templates also comes with a list of boring extensions ignored by default.

```yaml
# Custom extensions denylist
denylist:
  - .avi
  - .mov
  - .mp3
```

The default nuclei denylist contains the following extensions - 

```
3g2,3gp,7z,apk,arj,avi,axd,bmp,css,csv,deb,dll,doc,drv,eot,exe,flv,gif,gifv,gz,h264,ico,iso,jar,jpeg,jpg,lock,m4a,m4v,map,mkv,mov,mp3,mp4,mpeg,mpg,msi,ogg,ogm,ogv,otf,pdf,pkg,png,ppt,psd,rar,rm,rpm,svg,swf,sys,tar,tar.gz,tif,tiff,ttf,txt,vob,wav,webm,wmv,woff,woff2,xcf,xls,xlsx,zip
```

## max-size

Maximum size of the file to perform matching on.

By default, nuclei will process 1 GB of content and not go more than that. It can be set to much lower or higher depending on use. If set to "no" then all content will be processed

```yaml
# Max size of 5 mb should be processed
max-size: 5Mb
```

## no-recursive

`no-recursive` option disables recursive walking of directories / globs while input is being processed for file module of nuclei.

```yaml
# do not recursively walk folders
no-recursive: true
```

## Example

The final example template file for a Private Key detection is provided below.

```yaml
id: google-api-key

info:
  name: Google API Key
  author: pdteam
  severity: info

file:
  - extensions:
      - all
      - txt

    extractors:
      - type: regex
        name: google-api-key
        regex:
          - "AIza[0-9A-Za-z\\-_]{35}"
```

```bash
# Running file template on http-response/ directory
nuclei -t file.yaml -target http-response/

# Running file template on output.txt
nuclei -t file.yaml -target output.txt
```
