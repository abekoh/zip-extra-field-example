from zipfile import ZipFile
import os
import time

os.environ["TZ"] = "UTC"
time.tzset()

with ZipFile("/tmp/with_python.zip", "w") as zipFile:
    zipFile.writestr("with_python.txt", "with_python")
