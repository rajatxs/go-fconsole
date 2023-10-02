import ImageTool from '@editorjs/image';
import {DeletePostImage} from '../../wailsjs/go/services/PostService';

export default class CustomImageTool extends ImageTool {
   async removed() {
      if (!this._data.file) {
         return;
      }

      const publicId = this._data.file.publicId;
      await DeletePostImage(publicId);
   }
}
