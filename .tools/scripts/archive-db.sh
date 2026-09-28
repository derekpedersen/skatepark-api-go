#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SOURCE_DIR="${ROOT_DIR}/.db"
BACKUP_DIR="${ROOT_DIR}/.old-db"
TIMESTAMP="$(date +"%Y%m%d-%H%M%S")"
ARCHIVE_NAME="skatepark-db-${TIMESTAMP}.zip"
ARCHIVE_PATH="${BACKUP_DIR}/${ARCHIVE_NAME}"

if [[ ! -d "${SOURCE_DIR}" ]]; then
	echo "source directory not found: ${SOURCE_DIR}" >&2
	exit 1
fi

mkdir -p "${BACKUP_DIR}"

(
	cd "${ROOT_DIR}"
	zip -rq "${ARCHIVE_PATH}" .db
)

echo "created backup archive: ${ARCHIVE_PATH}"