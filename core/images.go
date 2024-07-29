package core

import (
	"io"

	"github.com/GNUWeeb/asuka-core/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
)

func (a *Asuka) ImageList(opt image.ListOptions) (*[]image.Summary, error) {
	list, err := a.client.ImageList(a.ctx, opt)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (a *Asuka) ImagesHistory(image string) ([]image.HistoryResponseItem, error) {
	img, err := a.client.ImageHistory(a.ctx, image)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (a *Asuka) ImagesBuild(opt types.ImageBuildOptions, src, dockerFile string, f func(rd io.Reader) error) error {
	tar, err := utils.TarWithOpt(src)
	if err != nil {
		return err
	}

	res, err := a.client.ImageBuild(a.ctx, tar, opt)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Hook function
	return f(res.Body)
}

func (a *Asuka) ImagesPush(image string, opt image.PushOptions, f func(rd io.Reader) error) error {

	res, err := a.client.ImagePush(a.ctx, image, opt)
	if err != nil {
		return err
	}
	defer res.Close()

	// Hook function
	return f(res)
}

func (a *Asuka) ImagesPull(ref string, opt image.PullOptions, f func(rd io.Reader) error) error {

	res, err := a.client.ImagePull(a.ctx, ref, opt)
	if err != nil {
		return err
	}
	defer res.Close()

	// Hook function
	return f(res)
}

func (a *Asuka) ImagesImport(ref string, src image.ImportSource, opt image.ImportOptions, f func(rd io.Reader) error) error {

	res, err := a.client.ImageImport(a.ctx, src, ref, opt)
	if err != nil {
		return err
	}
	defer res.Close()

	// Hook function
	return f(res)
}

func (a *Asuka) ImagesInspect(image string) (*types.ImageInspect, []byte, error) {
	inspect, body, err := a.client.ImageInspectWithRaw(a.ctx, image)
	if err != nil {
		return nil, nil, err
	}

	return &inspect, body, nil
}

func (a *Asuka) ImagesRemove(image string, opt image.RemoveOptions) (*[]image.DeleteResponse, error) {

	res, err := a.client.ImageRemove(a.ctx, image, opt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *Asuka) ImagesSave(image []string, f func(rd io.Reader) error) error {

	res, err := a.client.ImageSave(a.ctx, image)
	if err != nil {
		return err
	}

	// Hook function
	return f(res)
}

func (a *Asuka) ImagesSearch(keyword string, opt registry.SearchOptions) (*[]registry.SearchResult, error) {

	res, err := a.client.ImageSearch(a.ctx, keyword, opt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *Asuka) ImagesTag(src, dest string) error {

	err := a.client.ImageTag(a.ctx, src, dest)
	if err != nil {
		return err
	}

	return nil
}
