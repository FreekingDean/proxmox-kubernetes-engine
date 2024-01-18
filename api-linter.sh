set -eu

temp_dir="$(mktemp -d)"
descriptors_path=$temp_dir/descriptors.textpb
buf build -o "${descriptors_path}" --as-file-descriptor-set

cd protos
all=$(find -name *.proto)
api-linter \
    --descriptor-set-in "${descriptors_path}" \
    --config ../.api-linter.yml \
    ${all}

rm -r "${temp_dir}"
