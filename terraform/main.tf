provider "google" {
  project = var.project
  region  = var.region
}

module "function_GetAlbumById"{
  source      = "./modules/function"
  source_dir  = "../functions"
  project     = var.project
  function_name = "albums_function"
  function_entry_point = "GetAlbums"
}

