src=""
for file in ./*.go; do
  if [[ $file == *_test.go ]]; then
    continue
  fi
  src="${src} ${file}"
done
go build -o anonchat $src