<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>
      <form method="post" action="http://localhost:8888/api/v1/vmiss/upload?key=folder" id="folder-uploader" enctype="multipart/form-data">
        <input type="file" name="folder" webkitdirectory>
        <input type="submit" value="Send" />
      </form>
    </div>
  </div>
</template>
